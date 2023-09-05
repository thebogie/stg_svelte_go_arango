import {GraphQLClient} from "graphql-request"
import {PUBLIC_API_BASE_URL} from '$env/static/public'
import type {IPlayer, IPlayerLogin, IPlayerPayload} from "$lib/interfaces/player.interface";
import {gql_loginuser} from "$lib/gql/player.js";
import type {CustomError} from '$lib/interfaces/error.interface';

export const loginUser = async (email: string, password: string): Promise<[object, Array<CustomError>]> => {

    var responseplayer : IPlayer = {
        email: email,
    }
   var response :IPlayerLogin = {
        token: "",
       userdata: responseplayer,
   };
    var errors : CustomError[] = [{
        error: ""
    }];

    const graphQLClient = new GraphQLClient(PUBLIC_API_BASE_URL, {
        credentials: 'include',
        mode: 'cors',

    })

    const input: IPlayerPayload =
        {
            username: email,
            password: password,
        }

    try {
        const res : object = await graphQLClient.request(gql_loginuser, {input}) as object;
        // Handle the successful response here
        response.userdata = res.loginUser.userdata
        response.token = res.loginUser.token


    } catch (err : any) {
        // Handle the error here

        if (err.response.status == 403) {
            console.error("GraphQL request error:", err);
            errors[0].error = responseplayer.email + " was unable to login. Forbidden. Wrong password? Isnt registered?";
        }


    }

    return [response, errors];
};


export const logOutUser = async (): Promise<void> => {

};