export interface IPlayer {
	_key?: string;
	_id?: string;
	_rev?: string;
	firstname?: string;
	email?: string;
	password?: string;
	accessToken?: string | undefined;
	roles?: Array<string>;
}

export interface ILoginUser {
	token: string;
	userdata: IPlayer;
}

//Player Nemesis
export interface INemesisLeaderboard {
	players: INemesisRecord[];
	nemesis: INemesisRecord[];
	punchingBag: INemesisRecord[];
	winnersTogether: INemesisRecord[];
	losersTogether: INemesisRecord[];


}

export interface INemesisRecord {
	player: IPlayer;
	winsAgainstYou: number;
	lostAgainstYou: number;
	winsTogether: number;
	lostTogether: number;
}
