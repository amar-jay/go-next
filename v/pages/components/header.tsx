import Link from "next/link"
import { signIn, signOut, useSession } from "next-auth/react"
import styles from "./header.module.css"

interface Route {
  href: `/${string}`
  name: string
}
const routes:Route[] = [
  {
    name: "Home",
    href: "/"
  },
  {
    name: "Editor",
    href: "/editor"
  },
  {
    name: "Client",
    href: "/client"
  },
  {
    name: "Server",
    href: "/server"
  },
  {
    name: "Protected",
    href: "/protected"
  },
  {
    name: "API",
    href: "/api-example"
  },
  {
    name: "Admin",
    href: "/admin"
  },
  {
    name: "Me",
    href: "/me"
  }
]
// The approach used in this component shows how to build a sign in and sign out
// component that works on pages which support both client and server side
// rendering, and avoids any flash incorrect content on initial page load.
export default function Header() {
  const { data: session, status } = useSession()
  const loading = status === "loading"

  return (
    <header>
      <noscript>
        <style>{`.nojs-show { opacity: 1; top: 0; }`}</style>
      </noscript>
      <div className={styles.signedInStatus}>
        <p
          className={`nojs-show ${
            !session && loading ? styles.loading : styles.loaded
          }`}
        >
          {!session && (
            <>
              <span className={styles.notSignedInText}>
                You are not signed in
              </span>
              <a
                href={`/api/auth/signin`}
                className={styles.buttonPrimary}
                onClick={(e) => {
                  e.preventDefault()
                  signIn()
                }}
              >
                Sign in
              </a>
            </>
          )}
          {session?.user && (
            <>
              {session.user.image && (
                <span
                  style={{ backgroundImage: `url('${session.user.image}')` }}
                  className={styles.avatar}
                />
              )}
              <span className={styles.signedInText}>
                <small>Signed in as</small>
                <br />
                <strong>{session.user.email ?? session.user.name}</strong>
              </span>
              <a
                href={`/api/auth/signout`}
                className={styles.button}
                onClick={(e) => {
                  e.preventDefault()
                  signOut()
                }}
              >
                Sign out
              </a>
            </>
          )}
        </p>
      </div>
      <nav>
        <ul className={styles.navItems}>
	  {routes.map((e, idx) => (
	    <li key={idx} className={styles.navItem}>
	      <Link href={e.href}>{e.name}</Link>
	      {" | "}
	    </li>
	  ))}
        </ul>
      </nav>
    </header>
  )
}
