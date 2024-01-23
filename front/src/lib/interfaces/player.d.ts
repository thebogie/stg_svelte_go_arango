export interface IPlayer {
	_key?: string;
	_id?: string;
	rev?: string;
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

