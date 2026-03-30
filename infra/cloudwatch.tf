resource "aws_cloudwatch_log_group" "eks" {
  name              = "/aws/eks/${aws_eks_cluster.eks.name}/cluster"
  retention_in_days = 30
}