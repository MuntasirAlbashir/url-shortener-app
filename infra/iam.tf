
resource "aws_iam_openid_connect_provider" "hcp_terraform" {
  url             = "https://app.terraform.io"
  client_id_list  = ["aws.workload.identity"]
}


resource "aws_iam_role" "hcp_terraform" {
  name = "hcp-terraform-url-app"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Federated = aws_iam_openid_connect_provider.hcp_terraform.arn
      }
      Action = "sts:AssumeRoleWithWebIdentity"
      Condition = {
        StringEquals = {
          "app.terraform.io:aud" = "aws.workload.identity"
        }
        StringLike = {
          "app.terraform.io:sub" = "organization:${var.hcp_org}:project:*:workspace:${var.hcp_workspace}:run_phase:*"
        }
      }
    }]
  })
}

resource "aws_iam_role_policy_attachment" "hcp_terraform_admin" {
  role       = aws_iam_role.hcp_terraform.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}

output "hcp_terraform_role_arn" {
  value = aws_iam_role.hcp_terraform.arn
}
