import { _graphql } from '$lib/services/graphql';
import { gql } from 'graphql-request';
import type {IPlayer} from '$lib/interfaces/player.d';


const getContestByPlayer = async (token: string) => {
	// Check if user exists
	console.log('getContestByPlayer Service');
};
//export { createUser, loginUser };
export { getContestByPlayer };
