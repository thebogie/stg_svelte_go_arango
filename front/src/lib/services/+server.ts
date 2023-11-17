import { gql, GraphQLClient } from 'graphql-request';
import { envvars } from '$lib/utils/constrants';
import type { IPlayer } from '$lib/interfaces/player.interface';

export async function graphql(query: string, variables: any) {
	interface keyable {
		[key: string]: any;
	}

	let result: IPlayer = {};

	const graphQLClient = new GraphQLClient(`${envvars.BASE_API_URI}`, {
		credentials: 'include',
		mode: 'cors',
		headers: {
			//Authorization: `Bearer fish`
		}
	});

	try {
		let results = (await graphQLClient.request(query, variables)) as object;
		result.accessToken = results.loginUser.token;
		result.email = results.loginUser.userdata.email;
		result._key = results.loginUser.userdata._key;

		console.log('GRAPHQL:' + JSON.stringify(result));
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

	return result;
}
