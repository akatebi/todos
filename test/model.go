package model

const UserQuery string = `
query User($email: String!) {
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

type Variables struct {
	Email string `json:"email"`
}
type GraphQL struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}
