terraform {
  cloud {
    organization = "monty-demo"

    workspaces {
      name = "url-app"
    }
  }

  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}
