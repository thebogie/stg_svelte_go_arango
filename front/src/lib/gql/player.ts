import {gql} from "graphql-request"


export const gql_loginuser = gql`
    mutation LoginUser($input : Login!) {
        loginUser(input: $input) {
            token
            userdata {
                _key
                _id
                rev
                firstname
                email
                password
            }
        }
    }
`