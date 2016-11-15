provider "aws" {
    region = "ap-northeast-2"
}

resource "aws_key_pair" "deployer" {
  key_name = "deployer-key"
  public_key = "${file("id_rsa.pub")}"
}

resource "aws_instance" "web" {
    ami = "ami-a04297ce"
    instance_type = "t2.micro"
    key_name = "${aws_key_pair.deployer.key_name}"
    security_groups = ["${aws_security_group.allow_all.name}"]

    connection {
        type = "ssh"
        user = "ec2-user"
        private_key = "${file("id_rsa")}"
    }

    provisioner "file" {
        source = "klotter"
        destination = "/tmp/klotter"
    }
}

resource "aws_security_group" "allow_all" {
  name = "allow_all"
  description = "Allow all inbound traffic"

  ingress {
      from_port = 0
      to_port = 0
      protocol = "-1"
      cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
      from_port = 0
      to_port = 0
      protocol = "-1"
      cidr_blocks = ["0.0.0.0/0"]
  }
}