import { _graphql } from '$lib/services/graphql';
import { gql } from 'graphql-request';
import type { IPlayer } from '$lib/interfaces/player';
import type { IContest } from '$lib/interfaces/contest';
import Cookies from 'js-cookie';
import type { Graphql_variables } from '$lib/interfaces/graphql';

interface responseObject {
	[key: string]: never;
}

const loginPlayer = async (email: string, password: string) => {
	// Check if user exists
	console.log('LoginUser Service');

	let results: unknown;
	const player_found: IPlayer = {};

	const gql_loginuser = gql`
		mutation LoginUser($input: Login!) {
			LoginUser(input: $input) {
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
		results = await _graphql(player_found, gql_loginuser, {
			input: { username: email, password: password }
		});

		player_found.accessToken = results.LoginUser.token;
		player_found.email = results.LoginUser.userdata.email;
		player_found._key = results.LoginUser.userdata._key;

		if (player_found.accessToken) {
			Cookies.set('player', JSON.stringify(results.LoginUser.userdata), {
				expires: 1, // The cookie will expire in 7 days
				path: '/', // The cookie is valid for all paths on your domain
				secure: true, // This will make the cookie be sent only over HTTPS
				sameSite: 'strict' // This will avoid cookie being sent in cross-site requests;
			});
		}
	} catch (err: Error) {
		console.log('ERror: ' + JSON.stringify(err.message));
		throw err;
	}
	console.log('LOGINPLAYERSERVICE: ' + JSON.stringify(player_found));

	return player_found;
};

const getPlayerTotalResults = async (playerdata: IPlayer): Promise<IContest[]> => {
	// Check if user exists
	console.log('getPlayerTotalResults Service');
	let response: responseObject;

	const variables = {
		player: 'player/' + playerdata._key
	};
	const gql_query = gql`
		query GetContestsPlayerTotalResults($player: String!) {
			GetContestsPlayerTotalResults(player: $player) {
				_key
				_id
				outcomes {
					player
					place
					result
				}
			}
		}
	`;

	try {
		response = await _graphql(playerdata, gql_query, variables);
	} catch (err: any) {
		console.log('ERROR in getPlayerGamesPlayed:' + JSON.stringify(err.message));
		throw err;
	}

	// Now you can map your JSON data to the IContest interface
	const contests: IContest[] = response['GetContestsPlayerTotalResults'].map((contestData: any) => {
		return {
			_key: contestData._key,
			_id: contestData._id,
			outcomes: contestData.outcomes
		};
	});

	//console.log('getPlayerTotalResults: ' + JSON.stringify(contests));

	return contests;
};
//export { createUser, loginUser };
export { loginPlayer, getPlayerTotalResults };


/* const checkPlayer = async (playerdata : IPlayer) => {
	// Check if user exists
	console.log('CheckLogin Service');

	var token = '';
	var error: string = '';
	let results: any;
	let response: IPlayer = {};

	const variables = {
		player: playerdata.email
	};
	const gql_loginuser = gql`
		query CheckLogin($player: String!) {
			checklogin(player: $player) 
		}
	`;

	try {
		results = await _graphql(playerdata, gql_loginuser, variables);


	} catch (err: any) {
		console.log('ERror: ' + JSON.stringify(err.message));
		throw err;
	}
	console.log('CHECKLOGIN RESULTS: ' + JSON.stringify(results));

	return results.checklogin;
};

 */
