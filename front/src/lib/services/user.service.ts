import {graphql} from "$lib/services/+server";
import {gql} from "graphql-request"
import type {CustomError} from "$lib/interfaces/error.interface";

const loginUser = async (email: string, password: string) => {
    // Check if user exists
    console.log("LoginUser Service");

    var token = "";
    var error : string = "";
    var response;


const variables = {
        input: {"username":email,"password":password},
    };
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

    try {
        response = await graphql(gql_loginuser, variables);

    } catch (err:any) {
        console.log("ERror: " + JSON.stringify(err.message));
        throw err;
    }
    console.log("RESPONSE: " + JSON.stringify(response));
    return  token ;
}
//export { createUser, loginUser };
export {  loginUser };