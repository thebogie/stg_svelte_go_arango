import {GraphQLClient} from "graphql-request"
import {PUBLIC_API_BASE_URL} from '$env/static/public'
import type {IPlayerPayload} from "$lib/interfaces/player.interface";
import {gql_loginuser} from "$lib/gql/player.js";
import type {CustomError} from '$lib/interfaces/error.interface';

export const loginUser = async (email: string, password: string): Promise<[object, Array<CustomError>]> => {
    const graphQLClient = new GraphQLClient(PUBLIC_API_BASE_URL, {
        credentials: 'include',
        mode: 'cors',

    })

    const input: IPlayerPayload =
        {
            username: email,
            password: password,
        }


    const response : object = await graphQLClient.request(gql_loginuser, {input}) as object;
    // const response = await res.json();
    // console.log("RES" + JSON.stringify(response));

    return [response, []];

};


export const logOutUser = async (): Promise<void> => {

};