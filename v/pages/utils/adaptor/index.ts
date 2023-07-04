import { Adapter, DefaultAdapter } from "next-auth/adapters"
import { redirect } from "next/dist/server/api-utils";
import { NextResponse } from "next/server";


const request = (url: string) => {
	const req = fetch(url)
	.then(res => res.json())
	.then(r => {
		switch (r.status){
			case 200:
				return r;
			case 404:
        // redirect('/404')
				NextResponse.redirect('/404')
        return
			default:
				NextResponse.redirect('/error')
        return
		}
			return r
		
	})

	return req;
}
const A = (url: string): Adapter => {
  return {
    async createUser(user) {
		console.log(JSON.stringify(user))
      return {} as any
    },
    async getUser(id) {
		console.log(JSON.stringify(id))
      return {} as any
    },
    async getUserByEmail(email) {
		console.log(JSON.stringify(email))
      return {} as any
    },
    async getUserByAccount({ providerAccountId, provider, ...rest }) {
		console.log(JSON.stringify({providerAccountId, provider, ...rest}))
      return {} as any
    },
    async updateUser(user) {
		console.log(JSON.stringify(user))
      return {} as any
    },
    async deleteUser(userId) {
		console.log(JSON.stringify(userId))
      return
    },
    async linkAccount(account) {
		console.log(JSON.stringify(account))
      return
    },
    async unlinkAccount({ providerAccountId, provider, ...rest }) {
		console.log(JSON.stringify({providerAccountId, provider, ...rest}))
      return
    },
    async createSession({ sessionToken, userId, expires, ...rest }) {
		console.log(JSON.stringify({ sessionToken, userId, expires, ...rest}))
      return {} as any
    },
    async getSessionAndUser(sessionToken) {
		console.log(JSON.stringify(sessionToken))
      return {} as any
    },
    async updateSession({ sessionToken, ...rest }) {
		console.log(JSON.stringify({sessionToken, ...rest}))
      return {} as any
    },
    async deleteSession(sessionToken) {
		console.log(JSON.stringify(sessionToken))
      return
    },
    async createVerificationToken({ identifier, expires, token, ...rest }) {
		console.log(JSON.stringify({ identifier, expires, token, ...rest }))
      return {} as any
    },
    async useVerificationToken({ identifier, token, ...rest }) {
		console.log(JSON.stringify({ identifier, token, ...rest }))
      return {} as any
    },
  } 
}

export default A