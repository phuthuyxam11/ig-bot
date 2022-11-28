import {createRoot} from 'react-dom/client'
import {Chart, registerables} from 'chart.js'

// Apps
import {App} from './app/App'
import './_metronic/assets/sass/style.scss'
import 'bootstrap'
/**
 * Base URL of the website.
 *
 * @see https://facebook.github.io/create-react-app/docs/using-the-public-folder
 */
const {PUBLIC_URL} = process.env
Chart.register(...registerables)
const container = document.getElementById('root')
if (container) {
  createRoot(container).render(<App basename={PUBLIC_URL} />)
}
