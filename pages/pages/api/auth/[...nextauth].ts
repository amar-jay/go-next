import NextAuth, { NextAuthOptions } from "next-auth"
import GoogleProvider from "next-auth/providers/google"
import FacebookProvider from "next-auth/providers/facebook"
import GithubProvider from "next-auth/providers/github"
//import TwitterProvider from "next-auth/providers/twitter"
// import AppleProvider from "next-auth/providers/apple"
// import EmailProvider from "next-auth/providers/email"

// For more information on each option (and a full list of options) go to
// https://next-auth.js.org/configuration/options
export const authOptions = {
  // https://next-auth.js.org/configuration/providers/oauth
  providers: [
    /* EmailProvider({
         server: process.env.EMAIL_SERVER,
         from: process.env.EMAIL_FROM,
       }),
    // Temporarily removing the Apple provider from the demo site as the
    // callback URL for it needs updating due to Vercel changing domains

    Providers.Apple({
      clientId: process.env.APPLE_ID,
      clientSecret: {
        appleId: process.env.APPLE_ID,
        teamId: process.env.APPLE_TEAM_ID,
        privateKey: process.env.APPLE_PRIVATE_KEY,
        keyId: process.env.APPLE_KEY_ID,
      },
    }),
    */
    FacebookProvider({
      clientId: process.env.FACEBOOK_ID,
      clientSecret: process.env.FACEBOOK_SECRET,
    }),
    GithubProvider({
      clientId: process.env.GITHUB_ID,
      clientSecret: process.env.GITHUB_SECRET,
    }),
    GoogleProvider({
      clientId: process.env.GOOGLE_ID,
      clientSecret: process.env.GOOGLE_SECRET,
    }),
    // TwitterProvider({
    //   clientId: process.env.TWITTER_ID,
    //   clientSecret: process.env.TWITTER_SECRET,
    // })
  ],
  theme: {
    colorScheme: "light",
  },
  callbacks: {
    /*
    async signIn({user, account, profile}) {
      if (!profile) {
        return false 
      }
      if (account?.provider === "facebook" || "google" || "github") {
        user.name = profile.name
        user.email = profile.email
        user.image = profile.image
        return true
      }
      return false
    },
    */
    async jwt({ token, user }) {
      token.userRole = "admin"
      console.log("user from jwt: \n", JSON.stringify(user))

      return token
    },
    async session({ session, user}) {
      console.log("user from session: \n", JSON.stringify(user))
      return session
    }
  },
} satisfies NextAuthOptions;

export default NextAuth(authOptions)
