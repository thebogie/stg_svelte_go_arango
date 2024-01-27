import { GraphQLClient, type Variables } from 'graphql-request';
import type { IPlayer } from '$lib/interfaces/player';

export async function _graphql(playerdata: IPlayer, query: string, variables: Variables) {
	let results: unknown;

	let authCookie = playerdata.accessToken;
	//console.log('AUTH = ' + authCookie);

	if (!authCookie) {
		authCookie = '';
	}
	const apiUrl = import.meta.env.VITE_BASE_API_URI;
	console.log('apiURL' + apiUrl);
	const graphQLClient = new GraphQLClient(apiUrl, {
		credentials: 'include',
		mode: 'cors',
		headers: {
			Authorization: authCookie
		}
	});

	try {
		if (Object.keys(variables).length === 0) {
			results = (await graphQLClient.request(query)) as object;
		} else {
			results = (await graphQLClient.request(query, variables)) as object;
		}

		//console.log('GRAPHQL:' + JSON.stringify(results));

		// @ts-expect-error this is error object
	} catch (err: Error) {
		// Handle the error here

		let username = (variables.input as { username: string }).username; // Correct assertion
		if (username === undefined) {
			username = 'unknown';
		}
		throw new Error(
			'ERR:' + username + ' was unable to login. Forbidden. Wrong password? Isnt registered?'
		);
	}
	//Max-Age=86400; Path=/; HttpOnly=true; sameSite=lax
	return results;
}
