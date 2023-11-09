import {gql, GraphQLClient} from "graphql-request"
import type {CustomError} from '$lib/interfaces/error.interface';
import {envvars} from "$lib/utils/constrants";

/** @type {import('./$types').RequestHandler} */
export async function graphql(query: string, variables: any) {
    interface keyable {
        [key: string]: any
    }


    let result = {};

    const graphQLClient = new GraphQLClient(`${envvars.BASE_API_URI}`, {
        credentials: 'include',
        mode: 'cors',

    });

    try {

        result = await graphQLClient.request(query, variables) as object;

        console.log("RES" + JSON.stringify(result));


    } catch (err: any) {
        // Handle the error here

        if (err.response.status == 403) {

            let username = "unknown";
            if (variables?.input?.username) {
                username = variables.input.username;
            }
            throw new Error("ERR:" + username + " was unable to login. Forbidden. Wrong password? Isnt registered?");
        } else {

            console.error("General GraphQL request error:", err);
        }


    }

    return result;
}