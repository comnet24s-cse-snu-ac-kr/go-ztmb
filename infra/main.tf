provider "aws" {
  region = "ap-northeast-2"
}

resource "aws_vpc" "main" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "main-vpc"
  }
}

resource "aws_subnet" "subnet1" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "172.16.10.0/24"
  availability_zone = "ap-northeast-2a"

  tags = {
    Name = "subnet1"
  }
}

resource "aws_subnet" "subnet2" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = "172.16.20.0/24"
  availability_zone = "ap-northeast-2a"

  tags = {
    Name = "subnet2"
  }
}

resource "aws_security_group" "sg1" {
  name        = "sg1"
  description = "Security group 1"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["147.46.0.0/16"]
  }

  ingress {
    from_port   = 853
    to_port     = 853
    protocol    = "tcp"
    cidr_blocks = ["172.16.20.0/24"]
  }

  ingress {
    from_port   = 8
    to_port     = 0
    protocol    = "icmp"
    cidr_blocks = ["172.16.10.0/24", "172.16.20.0/24"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "sg1"
  }
}

resource "aws_security_group" "sg2" {
  name        = "sg2"
  description = "Security group 2"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["147.46.0.0/16"]
  }

  ingress {
    from_port   = 853
    to_port     = 853
    protocol    = "tcp"
    cidr_blocks = ["172.16.10.0/24"]
  }

  ingress {
    from_port   = 8
    to_port     = 0
    protocol    = "icmp"
    cidr_blocks = ["172.16.10.0/24", "172.16.20.0/24"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "sg2"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "main-igw"
  }
}

resource "aws_route_table" "public_rt" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "public-rt"
  }
}

resource "aws_route_table_association" "a" {
  subnet_id      = aws_subnet.subnet1.id
  route_table_id = aws_route_table.public_rt.id
}

resource "aws_route_table_association" "b" {
  subnet_id      = aws_subnet.subnet2.id
  route_table_id = aws_route_table.public_rt.id
}

resource "tls_private_key" "ssh" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "aws_key_pair" "ssh" {
  key_name   = "ssh"
  public_key = tls_private_key.ssh.public_key_openssh
}

output "private_key" {
  value     = tls_private_key.ssh.private_key_pem
  sensitive = true
}

resource "aws_network_interface" "nic1" {
  subnet_id       = aws_subnet.subnet1.id
  private_ips     = ["172.16.10.254"]
  security_groups = [aws_security_group.sg1.id]
  source_dest_check           = false

  tags = {
    Name = "nic1"
  }
}

resource "aws_network_interface" "nic2" {
  subnet_id       = aws_subnet.subnet2.id
  private_ips     = ["172.16.20.254"]
  security_groups = [aws_security_group.sg2.id]
  source_dest_check           = false

  tags = {
    Name = "nic2"
  }
}

resource "aws_instance" "middlebox" {
  ami                         = data.aws_ami.amzn.id
  instance_type               = "t2.micro"  # t2.medium
  key_name                    = aws_key_pair.ssh.key_name

  network_interface {
    device_index         = 0
    network_interface_id = aws_network_interface.nic1.id
  }

  network_interface {
    device_index         = 1
    network_interface_id = aws_network_interface.nic2.id
  }

  user_data = <<-EOF
            #!/bin/bash
            hostnamectl set-hostname middlebox
            EOF

  tags = {
    Name = "middlebox"
  }
}

resource "aws_eip" "middlebox" {
  domain                    = "vpc"
  network_interface         = aws_network_interface.nic1.id
  associate_with_private_ip = "172.16.10.254"
}

# ---

resource "aws_instance" "bot" {
  ami                         = data.aws_ami.amzn.id
  instance_type               = "t2.micro"  # t3a.2xlarge
  subnet_id                   = aws_subnet.subnet1.id
  security_groups             = [aws_security_group.sg1.id]
  private_ip                  = "172.16.10.10"
  key_name                    = aws_key_pair.ssh.key_name

  user_data = <<-EOF
            #!/bin/bash
            ip route add 172.16.20.0/24 via 172.16.10.254
            hostnamectl set-hostname bot
            EOF

  tags = {
    Name = "bot"
  }
}

resource "aws_eip" "bot" {
  domain                    = "vpc"
  instance = aws_instance.bot.id
}

# ---

resource "aws_instance" "attacker" {
  ami                         = data.aws_ami.amzn.id
  instance_type               = "t2.micro"  # t2.medium
  subnet_id                   = aws_subnet.subnet2.id
  security_groups             = [aws_security_group.sg2.id]
  private_ip                  = "172.16.20.10"
  key_name                    = aws_key_pair.ssh.key_name

  user_data = <<-EOF
            #!/bin/bash
            ip route add 172.16.10.0/24 via 172.16.20.254
            hostnamectl set-hostname attacker
            EOF

  tags = {
    Name = "attacker"
  }
}

resource "aws_eip" "attacker" {
  domain                    = "vpc"
  instance = aws_instance.attacker.id
}
