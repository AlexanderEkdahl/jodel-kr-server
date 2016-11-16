variable "environment" {
  type = "string"
}

provider "aws" {
    region = "ap-northeast-2"
}

resource "aws_db_instance" "default" {
  apply_immediately    = true
  allocated_storage    = 10
  engine               = "postgres"
  engine_version       = "9.5.4"
  instance_class       = "db.t2.micro"
  name                 = "postgres"
  username             = "foo"
  password             = "barbarbarbar"
  publicly_accessible  = true
  // psql --host=tf-20161115075050596121757vhh.cftsw4tcoglb.ap-northeast-2.rds.amazonaws.com --port=5432 --username=foo --password --dbname=postgres
  // postgres://foo:barbarbarbar@tf-20161115075050596121757vhh.cftsw4tcoglb.ap-northeast-2.rds.amazonaws.com:5432/postgres
}

resource "aws_key_pair" "main" {
  key_name = "klottr-${var.environment}"
  public_key = "${file("id_rsa.pub")}"
}

resource "aws_instance" "web" {
    ami = "ami-983ce8f6"
    instance_type = "t2.micro"
    key_name = "${aws_key_pair.main.key_name}"
    security_groups = ["${aws_security_group.main.name}"]

    connection {
        type = "ssh"
        user = "ec2-user"
        private_key = "${file("id_rsa")}"
    }

    provisioner "file" {
        source = "../bin/klottr-linux-amd64"
        destination = "/tmp/klottr"
    }

    provisioner "file" {
        source = "../client/build"
        destination = "/tmp"
    }

    provisioner "remote-exec" {
        inline = [
            "mv /tmp/klottr $HOME/klottr",
            "chmod +x klottr",
            "echo postgres://${aws_db_instance.default.username}:${aws_db_instance.default.password}@${aws_db_instance.default.endpoint}/${aws_db_instance.default.name} > DATABASE_URL",
            "curl https://getcaddy.com | bash -s cors",
            "sudo setcap CAP_NET_BIND_SERVICE=+eip klottr",
            "sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/caddy",
            "mv /tmp/build $HOME/www"
        ]
    }
}

resource "aws_security_group" "main" {
  name = "klottr-${var.environment}"

  ingress {
      from_port = 22
      to_port = 22
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
      from_port = 80
      to_port = 80
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
      from_port = 443
      to_port = 443
      protocol = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
      from_port = 0
      to_port = 0
      protocol = "-1"
      cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_eip" "web" {
  instance = "${aws_instance.web.id}"
}

// data "template_file" "policy" {
//   template = "${file("${path.module}/policy.json")}"

//   vars {
//     bucket_name = "klottr-${var.environment}"
//   }
// }

// resource "aws_s3_bucket" "web" {
//     force_destroy = true
//     bucket = "klottr-${var.environment}"
//     acl = "public-read"
//     policy = "${data.template_file.policy.rendered}"

//     website {
//         index_document = "index.html"
//     }
// }

output "ip" {
    value = "${aws_eip.web.public_ip}"
}