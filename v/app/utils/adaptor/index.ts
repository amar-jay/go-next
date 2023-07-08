import { Adaptor } from "next-auth/adapters"
import { NextResponse } from "next/server"


const request = (url: string) => {
	const req = fetch(url)
	.then(res => res.json())
	.then(r => {
		switch (r.status){
			case 200:
				return r;
			case 404:
				NextResponse.redirect('/404')
			default:
				NextResponse.redirect('/error')
		}
			return r
		
	})

	return req;
}
/** @return { import("next-auth/adapters").Adapter } */
export default function Adapter(client, options = {}):Adaptor {
  return {
    async createUser(user) {
		console.log(JSON.stringify(user))
      return
    },
    async getUser(id) {
		console.log(JSON.stringify(id))
      return
    },
    async getUserByEmail(email) {
		console.log(JSON.stringify(email))
      return
    },
    async getUserByAccount({ providerAccountId, provider, ...rest }) {
		console.log(JSON.stringify({providerAccountId, provider, ...rest}))
      return
    },
    async updateUser(user) {
		console.log(JSON.stringify(user))
      return
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
      return
    },
    async getSessionAndUser(sessionToken) {
		console.log(JSON.stringify(sessionToken))
      return
    },
    async updateSession({ sessionToken, ...rest }) {
      return
		console.log(JSON.stringify({sessionToken, ...rest}))
    },
    async deleteSession(sessionToken) {
		console.log(JSON.stringify(sessionToken))
      return
    },
    async createVerificationToken({ identifier, expires, token, ...rest }) {
		console.log(JSON.stringify({ identifier, expires, token, ...rest }))
      return
    },
    async useVerificationToken({ identifier, token, ...rest }) {
		console.log(JSON.stringify({ identifier, token, ...rest }))
      return
    },
  }
}