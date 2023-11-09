import {gql, GraphQLClient} from "graphql-request"
import type {CustomError} from '$lib/interfaces/error.interface';
import { variables } from "$lib/utils/constrants";

/** @type {import('./$types').RequestHandler} */
export async function graphql(query: gql ) {

    var errors : CustomError[] = [];

    const graphQLClient = new GraphQLClient(`${variables.BASE_API_URI}`, {
        credentials: 'include',
        mode: 'cors',

    })

    const input =
        {
            username: "mitch@gmail.com",
            password: "letmein",
        }

    try {

        interface keyable {
            [key: string]: any
        }

        const res : keyable = await graphQLClient.request(query, {input}) as object;
        // Handle the successful response here

        //response.userdata = res.loginUser.userdata
        //response.token = res.loginUser.token


    } catch (err : any) {
        // Handle the error here

        if (err.response.status == 403) {
            console.error("GraphQL request error:", err);
            errors.push({error: "fish" + " was unable to login. Forbidden. Wrong password? Isnt registered?"});
        }


    }

    return [{}, errors];
}