import { Adapter, AdapterAccount, AdapterSession, AdapterUser, DefaultAdapter, VerificationToken } from "next-auth/adapters"
import { redirect } from "next/dist/server/api-utils";
import { NextResponse } from "next/server";
import path from "path";


const users: AdapterUser[] = []
var accounts: AdapterAccount[] = []
const sessions: AdapterSession[] = []
const removeAcc = (acc: {provider: string; providerAccountId: string}) => {
  const i = accounts.findIndex(a => a.provider === acc.provider && a.providerAccountId === acc.providerAccountId)
  accounts = [...accounts.slice(0, i), ...accounts.slice(i + 1)]
  return true
}

const getSession = async (sessionToken: string) => {
  const session = sessions.find(s => s.sessionToken === sessionToken)
  const user = users.find(u => u.id === session?.userId)
  return {session, user}
}

const getAcc = (acc: {provider: string; providerAccountId: string}) => {
  const account = accounts.find(a => a.provider === acc.provider && a.providerAccountId === acc.providerAccountId)
  const user = users.find(u => u.id === account?.userId)
  return user
}
const get = (id: string) => users.find(u => u.id === id)
const email = (email: string) => users.find(u => u.email === email)

const request = (p: string) => {
  if (!process.env.BACKEND_URL) throw new Error("BACKEND_URL is not defined")
  const url = path.join(process.env.BACKEND_URL, p)
	const req = fetch(url)
	.then(res => res.json())
	.then(r => {
		switch (r.status){
			case 200:
				return r.data;
			case 404:
        // redirect('/404')
				NextResponse.redirect('/404')
        return
			default:
        console.log("error occured in request: [url]", url, "[response]", r)
				NextResponse.redirect('/error')
        return
		}
	})

	return req;
}
const A = (url: string): Adapter => {
  return {
    async createUser(user) {
      const u = user as AdapterUser
      u.id = "1"
      users.push(u)
      console.log("from adapter[ createUser]", JSON.stringify(u))
      return u
    },
    async getUser(id) {
      const u = get(id)
      console.log("from adapter[ getUser ]", JSON.stringify(u))
      if (!u) return null
      return u
    },
    async getUserByEmail(e) {
      const u = email(e)
      console.log("from adapter[ getUserByEmail ]", JSON.stringify({...u, e}))
      if (!u) return null
      return u
    },
    async getUserByAccount({ providerAccountId, provider }) {
      const u = getAcc({ providerAccountId, provider })
    console.log("from adapter[ getUserByAccount ]", JSON.stringify({providerAccountId, provider, }), "\n[ user ]", JSON.stringify(u))
    if (!u) return null
    // TODO: for different providers
    // switch (provider){
    //   case "github":
    //     // const e = await request(`https://api.github.com/users/${providerAccountId}`)
    //     // providerAccountId = e.login
    //     return {
    //       email: "me@me.me",
    //       emailVerified: new Date(),
    //       name: "me",
    //       id: "1",
    //       image: "https://avatars.githubusercontent.com/u/1?v=4"
    //     } satisfies AdapterUser
    //   default:
    //     return u
    // }
    return u
    },
    async updateUser(user) {
      console.log("from adapter[ updateUser ]", JSON.stringify(user))
      return {
        ...user as Required<AdapterUser>,
        id: "1"
      }
    },
    async deleteUser(userId) {
      console.log("from adapter[ deleteUser ]", JSON.stringify(userId))
      return
    },
    async linkAccount(account) {
      accounts.push(account)
      console.log("from adapter[ linkAccount ]", JSON.stringify(account))
      return
    },
    async unlinkAccount({ providerAccountId, provider, ...rest }) {
      removeAcc({ providerAccountId, provider })
      console.log("from adapter[ unlinkAccount ]", JSON.stringify({providerAccountId, provider, ...rest}))
      return
    },
    async createSession({ sessionToken, userId, expires, ...rest }) {
      sessions.push({ sessionToken, userId, expires, ...rest })
    console.log("from adapter[ createSession ]", JSON.stringify({ sessionToken, userId, expires, ...rest}))
    return {
      sessionToken,
      userId,
      expires,
      ...rest
    }
    },
    async getSessionAndUser(sessionToken) {
     const {session, user} =  await getSession(sessionToken)
    console.log("from adapter[ getSessionAndUser ]", JSON.stringify(sessionToken), "\n[ session ]", JSON.stringify(session), "\n[ user ]", JSON.stringify(user))
    if (!session || !user) return null
      return {
        session,
        user
      }  
    },
    async updateSession({ sessionToken, ...rest }) {
      console.log("from adapter[ updateSession ]", JSON.stringify({ sessionToken, ...rest}))
      return {
        sessionToken,
        expires: new Date(),
        userId: "1",
      } satisfies AdapterSession
    },
    async deleteSession(sessionToken) {
      console.log("from adapter[ deleteSession ]", JSON.stringify(sessionToken))
      return {
        sessionToken,
        expires: new Date(),
        userId: "1",
      } satisfies AdapterSession
    },
    async createVerificationToken({ identifier, expires, token, ...rest }) {
      console.log("from adapter[ createVerificationToken ]", JSON.stringify({ identifier, expires, token, ...rest}))
      return {
        identifier,
        token,
        expires,
      } satisfies VerificationToken
    },
    async useVerificationToken({ identifier, token, ...rest }) {
      console.log("from adapter[ useVerificationToken ]", JSON.stringify({ identifier, token, ...rest}))
      return {
        identifier,
        token,
        expires: new Date(),
        ...rest,
      } satisfies VerificationToken
    },
  } 
}

export default A