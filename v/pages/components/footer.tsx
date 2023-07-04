import Link from "next/link"
import styles from "./footer.module.css"
import packageJSON from "../package.json"

const go_version = "1.19"
export default function Footer() {
  return (
    <footer className={styles.footer}>
      <hr />
      <ul className={styles.navItems}>
        <li className={styles.navItem}>
          <a href="https://links.themanan.me/go-next-docs">Documentation</a>
        </li>
        <li className={styles.navItem}>
          <a href="https://links.themanan.me/go-next-npm">NPM</a>
        </li>
        <li className={styles.navItem}>
          <a href="https://links.themanan.me/go-next">GitHub</a>
        </li>
        <li className={styles.navItem}>
          <em>next@{packageJSON.dependencies["next"]}</em> {' | '}
          <em>go@{go_version}</em>
        </li>
      </ul>
    </footer>
  )
}
