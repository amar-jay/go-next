import axios from 'axios';
import argon2 from 'argon2';
import { Adapter, AdapterAccount, AdapterSession, AdapterUser, DefaultAdapter, VerificationToken } from "next-auth/adapters"
import {nanoid} from 'nanoid';
const baseURL = process.env?.BACKEND_URL ?? 'http://localhost:4000/next';


const hash = async (id: string) => {

	let hashedUserId = argon2.hash(id, {
		type: argon2.argon2id,
		raw: true,
		secret: Buffer.from(process.env?.HASH_SECRET ?? "manan4real"),
		salt: Buffer.from(process.env?.HASH_SALT ?? "manan4real"),
	})
	return hashedUserId
}
const unhash = async (hashedUserId: string, userId: string) => {
	return argon2.verify(hashedUserId, userId, {
		type: argon2.argon2id,
		raw: true,
		secret: Buffer.from(process.env?.HASH_SECRET ?? "manan4real"),
		salt: Buffer.from(process.env?.HASH_SALT ?? "manan4real"),
	})
}
const request = axios.create({
	  baseURL,
	  timeout: 1000,
	  headers: {'X-Custom-Header': 'foobar'}
});

export const createUser = async (user: Omit<AdapterUser, "id">): Promise<AdapterUser> => {

	const id = nanoid()
	const {data} = await request.post('/create-user', {
		...user,
		id,
		// sid: await hash(id)
	})

	if (data.status !== 200) {
		throw new Error(data.message)
	}
  return {...user, id: data.data.id} 
}

export const getUser = async (id: string): Promise<AdapterUser | null> => {
	if (id.length < 1) {
		// throw new Error("id is empty")
		return null
	}

	const {data} = await request.get(`/get-user?id=${id}`)
	if (data.status !== 200) {
		// throw new Error(data.message)
		return null
	}
	return data.data as AdapterUser

}


export const getUserByEmail = async (email: string): Promise<AdapterUser | null> => {
	if (email.length < 1) {
		// throw new Error("email is empty")
		return null
	}

	const {data} = await request.get(`/get-user?email=${email}`)
	if (data.status !== 200) {
		// throw new Error(data.message)
		return null
	}
	return data.data as AdapterUser

}

export const getUserByAccount = async ({ providerAccountId, provider }: Pick<AdapterAccount, "provider" | "providerAccountId">): Promise<AdapterUser | null> => {
	const providers = ['github', 'google', 'facebook', 'email']
	if (providerAccountId.length < 1 || provider.length < 1) {
		// throw new Error("providerAccountId is empty")
		return null
	}

	if (providers.indexOf(provider) < 0) {
		// throw new Error("provider is not email")
		return null
	}
	const {data} = await request.get(`/get-user?provider=${provider}&providerAccountId=${providerAccountId}`)
	if (data.status !== 200) {
		// throw new Error(data.message)
		return null
	}
	return data.data as AdapterUser

}

export const updateUser = async (user: Partial<AdapterUser> & Pick<AdapterUser, "id">): Promise<AdapterUser> => {
	if (user.id.length < 1) {
		throw new Error("id is empty")
	}

	const {data} = await request.post('/update-user', user)
	if (data.status !== 200) {
		throw new Error(data.message)
	}
	return data.data as AdapterUser
}

export const deleteUser = async (id: string): Promise<void> => {
	const {data} = await request.post('/delete-user', {id})
	if (data.status !== 200) {
		throw new Error(data.message)
	}
}

export const linkAccount = async (account: AdapterAccount): Promise<void> => {
	// TODO: hash user id
	const {data} = await request.post('/link-account', {
		id: account.userId,
		sid: hash(account.userId),
		providerId: account.providerId,
		providerType: account.provider,
		providerAccountId: account.providerAccountId,
		refreshToken: account.refreshToken,
		accessToken: account.accessToken,
		accessTokenExpires: account.accessTokenExpires
	})
	if (data.status !== 200) {
		throw new Error(data.message)
	}
}

export const unlinkAccount = async ({provider, providerAccountId}:Pick<AdapterAccount, "provider" | "providerAccountId">): Promise<void> => {
	if (providerAccountId.length < 1 || provider.length < 1) {
		throw new Error("providerAccountId is empty")
	}

	const {data} = await request.post('/unlink-account', {
		provider,
		providerAccountId
	})

	if (data.status !== 200) {
		throw new Error(data.message)
	}
	return
}

export const createSession = async (session: AdapterSession): Promise<AdapterSession> => {
	const {userId, ...opts} = session
	const {data} = await request.post('/create-session', {
		id: userId,
		sid: hash(userId),
		...opts,
	})
	if (data.status !== 200) {
		throw new Error(data.message)
	}
	return data.data as AdapterSession
}

export const getSessionAndUser = async (sessionToken: string): Promise<{session: AdapterSession, user: AdapterUser}> => {
	if (sessionToken.length < 1) {
		throw new Error("sessionToken is empty")
	}

	const {data} = await request.get(`/get-session?token=${sessionToken}`)
	if (data.status !== 200) {
		throw new Error(data.message)
	}

	const session = data.data as AdapterSession
	const { data: user, status } = await request.get(`/get-user?id=${session.userId}`)
	if (status !== 200) {
		throw new Error(data.message)
	}

	// TODO: hash user id later
	// if (unhash(session.userId, user.id)) {
	// 	throw new Error("Invalid session")
	// }

	return {session, user: {
		...user,
	 }as AdapterUser}
}

export const updateSession = async (session: Partial<AdapterSession> & Pick<AdapterSession, "sessionToken">): Promise<AdapterSession> => {
	if (session.sessionToken.length < 1) {
		throw new Error("sessionToken is empty")
	}
	
	if (session.userId && session.userId.length < 1) {
		throw new Error("userId is empty")
	}

	const {data} = await request.post('/update-session', session)
	if (data.status !== 200) {
		throw new Error(data.message)
	}
	return data.data as AdapterSession
}

export const deleteSession = async (sessionToken: string): Promise<void> => {
	if (sessionToken.length < 1) {
		throw new Error("sessionToken is empty")
	}

	const {data} = await request.post(`/delete-session?token=${sessionToken}`)
	if (data.status !== 200) {
		throw new Error(data.message)
	}
}

export const createVerificationToken = async ({identifier, expires, token}: VerificationToken): Promise<VerificationToken> => {
	if (token.length < 1) {
		throw new Error("token is empty")
	}

	// if it is not expires, it will be expired in 24 hours
	if (expires < new Date()) {
		throw new Error("token is expired")
	}

	const {data} = await request.post('/create-verification-token', {
		identifier,
		expires,
		token
	})

	if (data.status !== 200) {
		throw new Error(data.message)
	}

	return {
		identifier: data.data.identifier,
		expires: data.data.expires,
		token: data.data.token
	}

}

export const useVerificationToken = async ({token, identifier}: { identifier: string, token: string}): Promise<VerificationToken> => {
	if (token.length < 1) {
		throw new Error("token is empty")
	}

	const {data} = await request.post('/use-verification-token', {token})
	if (data.status !== 200) {
		throw new Error(data.message)
	}

	return {
		identifier: data.data.identifier,
		expires: data.data.expires,
		token: data.data.token
	} as VerificationToken
}
