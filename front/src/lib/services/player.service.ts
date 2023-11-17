import { graphql } from '$lib/services/+server';
import { gql } from 'graphql-request';
import type { IPlayer } from '$lib/interfaces/player.interface';

const loginPlayer = async (email: string, password: string) => {
	// Check if user exists
	console.log('LoginUser Service');

	var token = '';
	var error: string = '';
	let results : any;
	let response: IPlayer = {};

	const variables = {
		input: { username: email, password: password }
	};
	const gql_loginuser = gql`
		mutation LoginUser($input: Login!) {
			loginUser(input: $input) {
				token
				userdata {
					_key
					_id
					rev
					firstname
					email
				}
			}
		}
	`;

	try {
		results = await graphql(token, gql_loginuser, variables) as object;
		// @ts-ignore
		response.accessToken = results.loginUser.token;
		// @ts-ignore
		response.email = results.loginUser.userdata.email;
		// @ts-ignore
		response._key = results.loginUser.userdata._key;
	} catch (err: any) {
		console.log('ERror: ' + JSON.stringify(err.message));
		throw err;
	}
	console.log('LOGINPLAYERSERVICE: ' + JSON.stringify(response));

	/* Set a cookie
    setCookie('token', response.loginUser.token, {
        path: '/',
        expires: new Date(Date.now() + 1000 * 60 * 60),
    }); */
	return response;
};

const getPlayerGamesPlayed = async (token: string) => {
	// Check if user exists
	console.log('getPlayerGamesPlayed Service');
	let response  = {};

	const gql_query = gql`
        query  Game {
            games { name }
        }
    `;

	try {
		response = await graphql(token, gql_query, "");

	} catch (err: any) {
		console.log('ERROR in getPlayerGamesPlayed:' + JSON.stringify(err.message));
		throw err;
	}
	console.log('GAMESERVICE: ' + JSON.stringify(response));

	return response;
};
//export { createUser, loginUser };
export { loginPlayer, getPlayerGamesPlayed };
