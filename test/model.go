package model

const userQuery string = `
query user($email: String!) {
	user(email: $email) {
	  id
	  email
	  completedCount
	  totalCount
	  todos(first: 1000) {
		edges {
		  cursor
		  node {
			id
			text
			complete
			__typename
		  }
		}
		pageInfo {
		  endCursor
		  hasNextPage
		}
	  }
	}
  `

type variables struct {
	email string
}

type body struct {
	query     string
	variables variables
}
