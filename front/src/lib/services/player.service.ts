import { _graphql } from '$lib/services/graphql';
import { gql } from 'graphql-request';
import type {
	ILoginUser,
	INemesisLeaderboard,
	INemesisRecord,
	IPlayer
} from '$lib/interfaces/player';
import type { IContest } from '$lib/interfaces/contest';
import Cookies from 'js-cookie';

/* loginPlayer: call api to login player and get token */
export const loginPlayer = async (email: string, password: string) => {
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
		player_found._id = _graphql_results.LoginUser.userdata._id;

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
export const getPlayerTotalResults = async (player_data: IPlayer): Promise<IContest[]> => {
	// Check if user exists
	console.log('getPlayerTotalResults Service');

	interface IGraphqlResults {
		GetContestsPlayerTotalResults: IContest[];
	}

	let contests: IContest[];

	const variables = {
		player: player_data._id
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
					player {
						_id
						firstname
						email
					}
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

/* createNemesisLeaderboard: get all the contests that a player played in */
export const createNemesisLeaderboard = async (
	player: IPlayer,
	contests: IContest[]
): Promise<INemesisLeaderboard> => {
	//let leaderBoardRecord : INemesisRecord = { };
	//let leaderBoard: INemesisLeaderboard = {players: leaderBoardRecord, currentStreak:1} ;
	console.log('createNemesisLeaderboard Service');

	const leaderBoard: INemesisLeaderboard = {
		players: [],
		nemesis: [],
		punchingBag: [],
		winnersTogether: [],
		losersTogether: []
	};

	contests.forEach(function (contest) {
		let myPlace = 0;
		contest.outcomes.forEach(function (outcome) {
			if (outcome.player._id == player._id) {
				//thats me!
				myPlace = outcome.place;
			}
		});

		contest.outcomes.forEach(function (outcome) {
			const enemy_id: string = outcome.player._id;

			//if this outcome is me... skip
			if (enemy_id === player._id) {
				return;
			}

			let enemy_record = leaderBoard.players.find((obj) => obj.player._id === enemy_id);

			if (!enemy_record) {
				leaderBoard.players.push({
					player: outcome.player,
					winsAgainstYou: 0,
					lostAgainstYou: 0,
					winsTogether: 0,
					lostTogether: 0
				});

				enemy_record = leaderBoard.players[leaderBoard.players.length - 1]; // Get the newly added object
			}

			if (outcome.place < myPlace) {
				if (outcome.result == 'won' || outcome.result == 'lost') {
					enemy_record.winsAgainstYou++;
				}
			}
			if (outcome.place > myPlace) {
				if (outcome.result == 'lost') {
					enemy_record.lostAgainstYou++;
				}
			}

			if (outcome.place == myPlace) {
				if (outcome.result == 'lost') {
					enemy_record.lostTogether++;
				}
				if (outcome.result == 'win') {
					enemy_record.winsTogether++;
				}
			}
		});
	});


	for (const playerRecord of leaderBoard.players) {
		if (playerRecord.winsAgainstYou !== 0) {
			leaderBoard.nemesis.push(playerRecord);
		}
		if (playerRecord.lostAgainstYou !== 0) {
			leaderBoard.punchingBag.push(playerRecord);
		}
		if (playerRecord.winsTogether !== 0) {
			leaderBoard.winnersTogether.push(playerRecord);
		}
		if (playerRecord.lostTogether !== 0) {
			leaderBoard.losersTogether.push(playerRecord);
		}
	}
	leaderBoard.nemesis = leaderBoard.nemesis.sort((n1, n2) => {
		if (n1.winsAgainstYou <= n2.winsAgainstYou) {
			return 1;
		}
		if (n1.winsAgainstYou > n2.winsAgainstYou) {
			return -1;
		}
		return 0;
	});
	leaderBoard.punchingBag = leaderBoard.punchingBag.sort((n1, n2) => {
		if (n1.lostAgainstYou <= n2.lostAgainstYou) {
			return 1;
		}
		if (n1.lostAgainstYou > n2.lostAgainstYou) {
			return -1;
		}
		return 0;
	});
	leaderBoard.winnersTogether = leaderBoard.winnersTogether.sort((n1, n2) => {
		if (n1.winsTogether <= n2.winsTogether) {
			return 1;
		}
		if (n1.winsTogether > n2.winsTogether) {
			return -1;
		}
		return 0;
	});
	leaderBoard.losersTogether = leaderBoard.losersTogether.sort((n1, n2) => {
		if (n1.lostTogether <= n2.lostTogether) {
			return 1;
		}
		if (n1.lostTogether > n2.lostTogether) {
			return -1;
		}
		return 0;
	});

	//	console.log('contests' + JSON.stringify(contests));

	return leaderBoard ;
};
