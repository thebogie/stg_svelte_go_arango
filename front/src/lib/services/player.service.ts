import { _graphql } from '$lib/services/graphql';
import { gql } from 'graphql-request';
import type { IPlayer } from '$lib/interfaces/player';
import type { IContest } from '$lib/interfaces/contest';
import Cookies from 'js-cookie';

interface responseObject {
	[key: string]: any;
}

const loginPlayer = async (email: string, password: string) => {
	// Check if user exists
	console.log('LoginUser Service');

	var token = '';
	var error: string = '';
	let results: any;
	let playerfound: IPlayer = {};

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
		results = await _graphql(playerfound, gql_loginuser, variables);

		// @ts-ignore
		playerfound.accessToken = results.loginUser.token;
		// @ts-ignore
		playerfound.email = results.loginUser.userdata.email;
		// @ts-ignore
		playerfound._key = results.loginUser.userdata._key;

		if (playerfound.accessToken) {
			Cookies.set('player', JSON.stringify(results.loginUser.userdata), {
				expires: 1, // The cookie will expire in 7 days
				path: '/', // The cookie is valid for all paths on your domain
				secure: true, // This will make the cookie be sent only over HTTPS
				sameSite: 'strict' // This will avoid cookie being sent in cross-site requests);
			});
		}
	} catch (err: any) {
		console.log('ERror: ' + JSON.stringify(err.message));
		throw err;
	}
	console.log('LOGINPLAYERSERVICE: ' + JSON.stringify(playerfound));

	return playerfound;
};

const getPlayerTotalResults = async (playerdata: IPlayer): Promise<IContest[]> => {
	// Check if user exists
	console.log('getPlayerTotalResults Service');
	var response: responseObject;

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
