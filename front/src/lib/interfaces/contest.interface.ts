import type {IGame} from "$lib/interfaces/game.interface";
import type {IVenue} from "$lib/interfaces/venue.interface";
import type {IOutcome} from "$lib/interfaces/outcome.interface";

export interface IContest {
	_key?: string;
	_id?: string;
	rev?: string;
	start?: string;
	startoffset?: string;
	stop?: string;
	stopoffset?: string;
	outcomes: IOutcome[];
	games: IGame[];
	venue: IVenue;
}



