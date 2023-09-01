export interface IPlayer {
    _key?: string;
    _id?: string;
    rev?: string;
    firstname?: string;
    email: string;
    password: string;
    accessToken: string | undefined;
    roles?: Array<string>;
}

export interface IPlayerPayload {
    username: string;
    password: string;
}

export interface IPlayerLogin {
    token: string;
    userdata: IPlayer;

}