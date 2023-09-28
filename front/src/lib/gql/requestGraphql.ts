import {GraphQLClient} from "graphql-request"
import type {IPlayer, IPlayerLogin, IPlayerPayload} from "$lib/interfaces/player.interface";
import {gql_loginuser} from "$lib/gql/player.js";
import type {CustomError} from '$lib/interfaces/error.interface';
import { variables } from "$lib/utils/constrants";

export const loginUser = async (email: string, password: string): Promise<[object : IPlayerLogin, Array<CustomError>]> => {

    var responseplayer : IPlayer = {
        email: email,
        password: "",
        accessToken: "",
    }
   var response :IPlayerLogin = {
        token: "",
       userdata: responseplayer,
   };
    var errors : CustomError[] = [];

    const graphQLClient = new GraphQLClient(`${variables.BASE_API_URI}`, {
        credentials: 'include',
        mode: 'cors',

    })

    const input: IPlayerPayload =
        {
            username: email,
            password: password,
        }

    try {

        interface keyable {
            [key: string]: any
        }

        const res : keyable = await graphQLClient.request(gql_loginuser, {input}) as object;
        // Handle the successful response here

        response.userdata = res.loginUser.userdata
        response.token = res.loginUser.token


    } catch (err : any) {
        // Handle the error here

        if (err.response.status == 403) {
            console.error("GraphQL request error:", err);
            errors.push({error: responseplayer.email + " was unable to login. Forbidden. Wrong password? Isnt registered?"});
        }


    }

    return [response, errors];
};


export const logOutUser = async (): Promise<void> => {

};