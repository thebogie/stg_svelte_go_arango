import type {IPlayer} from '$lib/interfaces/player';


export interface IOutcome {
	_key?: string;
	_id?: string;
	_rev?: string;
	player?: IPlayer;
	place?: number;
	result?: string;
}
