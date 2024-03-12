import type { IGame } from '@lib/interfaces/game.interface';
import type { IVenue } from '@lib/interfaces/venue.interface';
import type { IOutcome } from '@lib/interfaces/outcome.interface';
import type { INemesisLeaderboard, IPlayer } from '$lib/interfaces/player';

export interface DateRange {
	startDate: Date;
	endDate: Date;
}

export interface IContest {
	_key?: string;
	_id?: string;
	_rev?: string;
	start?: string;
	startoffset?: string;
	stop?: string;
	stopoffset?: string;
	outcomes: IOutcome[];
	games?: IGame[];
	venue?: IVenue;
}

export interface IProfile {
	total_results: IContest[];
	signed_in: IPlayer;
	personal_leader_board: INemesisLeaderboard;
}
