# tf-vars-git-committer

## Goal

Create a collection of GoLang modules, which would help reduce toil for certian DevOps tasks.

### GitBot Module

Acts as a developer and creates pull requests to make changes to structured code. The following languages are configured
* Terraform HCL

Given a Terraform repository, `gitbot` performs 

1. checkout
2. create a new branch with configured branch name pattern
3. updates to configured variables in `.tfvar` files
2. commit and push changes
3. create a PR 

### SlackBot Module 
