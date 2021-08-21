package model

const query string = `
query user($email: String!, $status: Status, $first: Int, $after: String, $last: Int, $before: String) {
	user(email: $email) {
	  id
	  email
	  completedCount
	  totalCount
	  todos(status: $status, first: $first, after: $after, last: $last, before: $before) {
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

type UserInput struct {
	email         string
	status        string
	first, last   string
	before, after string
}

type UserParams struct {
	query string
	variables UserInput 
}

type UserOutput struct {
	Data  interface{}
	Error interface{}
}

func QueryUser(userInput *UserInput) (UserOutput, error) {
	body := &UserParams{query: query, variables: userInput}}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	bodyBytes, err := graphql.Fetch(payloadBuf)
	var userOutput UserOutput
	json.Unmarshal(bodyBytes, &userOutput)
	fmt.Printf("%#v\n", userOutput)
	return userOutput, err
}
