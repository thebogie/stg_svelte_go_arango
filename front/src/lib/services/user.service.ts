import {graphql} from "$lib/services/+server";
import {gql} from "graphql-request"

const loginUser = async (email: string, password: string) => {
    // Check if user exists
    console.log("LoginUser Service");

     const gql_loginuser = gql`
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

let token = "toad";
let error : string = "fish";

    const response = graphql(gql_loginuser);
    return {error, token};
}
//export { createUser, loginUser };
export {  loginUser };