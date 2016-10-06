provider "aws" {
    region = "us-east-1"
}

data "aws_ami" "ubuntu" {
  most_recent = true
  filter {
    name = "name"
    values = ["ubuntu/images/ebs/ubuntu-trusty-14.04-amd64-server-*"]
  }
  filter {
    name = "virtualization-type"
    values = ["paravirtual"]
  }
  owners = ["099720109477"]
}

resource "aws_instance" "web" {
    ami = "${data.aws_ami.ubuntu.id}"
    instance_type = "t1.micro"

    // download s3 bucket binary 

    // restart if binary changes?

    // acl which allows access to binary

    // alternatively it creates
    // create what?!
    // argh
    // memory
}

// s3 bucket for binary

// s3 bucket object for binary
