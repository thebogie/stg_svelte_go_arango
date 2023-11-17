import { gql, GraphQLClient } from 'graphql-request';
import { envvars } from '$lib/utils/constrants';
import type { IPlayer } from '$lib/interfaces/player.interface';

export async function graphql(token: string, query: string, variables: any) {
	interface keyable {
		[key: string]: any;
	}

	let results: any = {};
	const graphQLClient = new GraphQLClient(`${envvars.BASE_API_URI}`, {
		credentials: 'include',
		mode: 'cors',
		headers: {
			Authorization: token
		}
	});

	try {
		if (variables == "") {
			results = (await graphQLClient.request(query)) as object;
		} else {
			results = (await graphQLClient.request(query, variables)) as object;
		}


		console.log('GRAPHQL:' + JSON.stringify(results));
	} catch (err: any) {
		// Handle the error here

		if (err.response.status == 403) {
			let username = 'unknown';
			if (variables?.input?.username) {
				username = variables.input.username;
			}
			throw new Error(
				'ERR:' + username + ' was unable to login. Forbidden. Wrong password? Isnt registered?'
			);
		} else {
			console.error('General GraphQL request error:', err);
		}
	}

	return results;
}
