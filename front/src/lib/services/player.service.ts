import { _graphql } from '$lib/services/graphql';
import { gql } from 'graphql-request';
import type { ILoginUser, IPlayer } from '$lib/interfaces/player';
import type { IContest } from '$lib/interfaces/contest';
import Cookies from 'js-cookie';

/* loginPlayer: call api to login player and get token */
const loginPlayer = async (email: string, password: string) => {
	// Check if user exists
	console.log('LoginUser Service');

	interface IGraphqlResults {
		LoginUser: ILoginUser;
	}

	//let _graphql_results: IGraphqlResults;

	const player_found: IPlayer = {};

	const gql_loginuser = gql`
		mutation LoginUser($input: Login!) {
			LoginUser(input: $input) {
				token
				userdata {
					_key
					_id
					firstname
					email
				}
			}
		}
	`;

	try {
		const _graphql_results = (await _graphql(player_found, gql_loginuser, {
			input: { username: email, password: password }
		})) as IGraphqlResults;

		player_found.accessToken = _graphql_results.LoginUser.token;
		player_found.email = _graphql_results.LoginUser.userdata.email;
		player_found._key = _graphql_results.LoginUser.userdata._key;

		if (player_found.accessToken) {
			Cookies.set('player', JSON.stringify(player_found), {
				expires: 1, // The cookie will expire in 7 days
				path: '/', // The cookie is valid for all paths on your domain
				secure: true, // This will make the cookie be sent only over HTTPS
				sameSite: 'strict' // This will avoid cookie being sent in cross-site requests;
			});
		}
	} catch (err) {
		if (err instanceof Error) {
			console.log('ERror: ' + JSON.stringify(err.message));
			throw err;
		}
	}
	console.log('LOGINPLAYERSERVICE: ' + JSON.stringify(player_found));

	return player_found;
};

/* getPlayerTotalResults: get all the contests that a player played in */
const getPlayerTotalResults = async (player_data: IPlayer): Promise<IContest[]> => {
	// Check if user exists
	console.log('getPlayerTotalResults Service');

	interface IGraphqlResults {
		GetContestsPlayerTotalResults: IContest[];
	}

	let contests: IContest[];

	const variables = {
		player: 'player/' + player_data._key
	};
	const gql_query = gql`
      query GetContestsPlayerTotalResults($player: String!) {
          GetContestsPlayerTotalResults(player: $player) {
              _id
              start
              startoffset
              stop
              stopoffset
              outcomes {
                  _id
                  place
                  result
                  player
              }
              games {
                  _id
                  name
              }
              venue {
                  _id
								address
								
              }
          }
      }
  `;

	try {
		const response = (await _graphql(player_data, gql_query, variables)) as IGraphqlResults;

		contests = response['GetContestsPlayerTotalResults'];
		/*
		// Now you can map your JSON data to the IContest interface
		contests = response['GetContestsPlayerTotalResults'].map((contestData) => {
			return {
				_key: contestData._key,
				_id: contestData._id,
				outcomes: contestData.outcomes
			};
		}); */
	} catch (err) {
		if (err instanceof Error) {
			console.log('ERROR in getPlayerGamesPlayed:' + JSON.stringify(err.message));
		}
		throw err;
	}

	//console.log('getPlayerTotalResults: ' + JSON.stringify(contests));

	return contests;
};

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
