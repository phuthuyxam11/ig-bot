import {Route, Navigate, Routes} from 'react-router-dom'
import {PageLink, PageTitle} from '../_metronic/layout/core'
import {Overview} from './gettings-started/Overview'
import {Accordion} from './base/Accordion'
import {Alerts} from './base/Alerts'
import {Badges} from './base/Badges'
import {Breadcrumb} from './base/Breadcrumb'
import {Bullets} from './base/Bullets'
import {Buttons} from './base/buttons/Buttons'
import {Cards} from './base/Cards'
import {Carousel} from './base/Carousel'
import {Forms} from './base/Forms'
import {HelpersBackground} from './base/HelpersBackground'
import {HelpersBorders} from './base/HelpersBorders'
import {HelpersFlexLayout} from './base/HelpersFlexLayout'
import {HelpersOpacity} from './base/HelpersOpacity'
import {HelpersText} from './base/HelpersText'
import {Indicator} from './base/Indicator'
import {Modal} from './base/Modal'
import {Overlay} from './base/Overlay'
import {Pagination} from './base/Pagination'
import {Pulse} from './base/Pulse'
import {Rating} from './base/Rating'
import {Rotate} from './base/Rotate'
import {Ribbon} from './base/Ribbon'
import {Separator} from './base/Separator'
import {Symbol} from './base/Symbol'
import {Tables} from './base/Tables'
import {Tabs} from './base/Tabs'
import {Utilities} from './base/utilities/Utilities'
import {NouiSlider} from './base/NouiSlider'
import {BootstrapIcons} from './icons/BootstrapIcons'
import {DuotuneIcons} from './icons/DuotuneIcons'
// import { StockholmIcons } from "./icons/StockholmIcons";
import {FontAwesomeIcons} from './icons/FontAwesomeIcons'
import {Changelog} from './gettings-started/Changelog'
import {CreatePage} from './gettings-started/CreatePage'
import {CRUDTables} from './gettings-started/CRUDTables'
import {Deployment} from './gettings-started/Deployment'
import {Internationalization} from './gettings-started/Internationalization'
import {References} from './gettings-started/References'
import {RTL} from './gettings-started/RTL'
import {SASSCustomization} from './gettings-started/SASSCustomization'
import {Skeleton} from './gettings-started/Skeleton'
import {DarkMode} from './gettings-started/DarkMode'
import {Updates} from './gettings-started/Updates'
import {QuickStart} from './gettings-started/QuickStart'
import {LineAwesomeIcons} from './icons/LineAwesomeIcons'
import {Backend} from './gettings-started/Backend'

const gettingStartedBC: Array<PageLink> = [
  {
    title: 'Getting started',
    isSeparator: false,
  },
  {
    title: '',
    isSeparator: true,
  },
]

const baseBC: Array<PageLink> = [
  {
    title: 'Base',
    isSeparator: false,
  },
  {
    title: '',
    isSeparator: true,
  },
]

const baseHelpersBC: Array<PageLink> = [
  {
    title: 'Base',
    isSeparator: false,
  },
  {
    title: '',
    isSeparator: true,
  },
  {
    title: 'Helpers',
    isSeparator: false,
  },
  {
    title: '',
    isSeparator: true,
  },
]

const iconsBC: Array<PageLink> = [
  {
    title: 'Icons',
    isSeparator: false,
  },
  {
    title: '',
    isSeparator: true,
  },
]

const DocsPage = () => (
  <Routes>
    <Route
      path='/docs/overview'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Overview</PageTitle>
          <Overview />
        </>
      }
    />

    <Route
      path='/docs/changelog'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Changelog</PageTitle>
          <Changelog />
        </>
      }
    />
    <Route
      path='/docs/create-a-page'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Create a page</PageTitle>
          <CreatePage />
        </>
      }
    />
    <Route
      path='/docs/deployment'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Deployment</PageTitle>
          <Deployment />
        </>
      }
    />
    <Route
      path='/docs/i18n'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Internationalization (i18n)</PageTitle>
          <Internationalization />
        </>
      }
    />
    <Route
      path='/docs/quick-start'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Quick Start</PageTitle>
          <QuickStart />
        </>
      }
    />
    <Route
      path='/docs/backend'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Back-end</PageTitle>
          <Backend />
        </>
      }
    />
    <Route
      path='/docs/references'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>References</PageTitle>
          <References />
        </>
      }
    />
    <Route
      path='/docs/dark-mode'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Dark Mode</PageTitle>
          <DarkMode />
        </>
      }
    />
    <Route
      path='/docs/updates'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Updates</PageTitle>
          <Updates />
        </>
      }
    />
    <Route
      path='/docs/rtl'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>RTL</PageTitle>
          <RTL />
        </>
      }
    />
    <Route
      path='/docs/sass-customization'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>SASS Customization</PageTitle>
          <SASSCustomization />
        </>
      }
    />
    <Route
      path='/docs/skeleton'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>Skeleton</PageTitle>
          <Skeleton />
        </>
      }
    />
    <Route
      path='/docs/crud'
      element={
        <>
          <PageTitle breadcrumbs={gettingStartedBC}>CRUD + Tables</PageTitle>
          <CRUDTables />
        </>
      }
    />

    <Route
      path='/docs/accordion'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Accordion</PageTitle>
          <Accordion />
        </>
      }
    />
    <Route
      path='/docs/alerts'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Alerts</PageTitle>
          <Alerts />
        </>
      }
    />
    <Route
      path='/docs/badges'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Basges</PageTitle>
          <Badges />
        </>
      }
    />
    <Route
      path='/docs/breadcrumb'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Breadcrumb</PageTitle>
          <Breadcrumb />
        </>
      }
    />
    <Route
      path='/docs/bullets'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Bullets</PageTitle>
          <Bullets />
        </>
      }
    />
    <Route
      path='/docs/buttons'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Buttons</PageTitle>
          <Buttons />
        </>
      }
    />
    <Route
      path='/docs/cards'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Cards</PageTitle>
          <Cards />
        </>
      }
    />
    <Route
      path='/docs/carousel'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Carousel</PageTitle>
          <Carousel />
        </>
      }
    />
    <Route
      path='/docs/forms'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Forms</PageTitle>
          <Forms />
        </>
      }
    />
    <Route
      path='/docs/helpers/background'
      element={
        <>
          <PageTitle breadcrumbs={baseHelpersBC}>Background</PageTitle>
          <HelpersBackground />
        </>
      }
    />
    <Route
      path='/docs/helpers/borders'
      element={
        <>
          <PageTitle breadcrumbs={baseHelpersBC}>Borders</PageTitle>
          <HelpersBorders />
        </>
      }
    />
    <Route
      path='/docs/helpers/flex-layout'
      element={
        <>
          <PageTitle breadcrumbs={baseHelpersBC}>Flex Layout</PageTitle>
          <HelpersFlexLayout />
        </>
      }
    />
    <Route
      path='/docs/helpers/opacity'
      element={
        <>
          <PageTitle breadcrumbs={baseHelpersBC}>Opacity</PageTitle>
          <HelpersOpacity />
        </>
      }
    />
    <Route
      path='/docs/helpers/text'
      element={
        <>
          <PageTitle breadcrumbs={baseHelpersBC}>Text</PageTitle>
          <HelpersText />
        </>
      }
    />
    <Route
      path='/docs/indicator'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Indicator</PageTitle>
          <Indicator />
        </>
      }
    />
    <Route
      path='/docs/modal'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Modal</PageTitle>
          <Modal />
        </>
      }
    />
    <Route
      path='/docs/overlay'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Overlay</PageTitle>
          <Overlay />
        </>
      }
    />
    <Route
      path='/docs/pagination'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Pagination</PageTitle>
          <Pagination />
        </>
      }
    />
    <Route
      path='/docs/pulse'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Pulse</PageTitle>
          <Pulse />
        </>
      }
    />
    <Route
      path='/docs/rating'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Rating</PageTitle>
          <Rating />
        </>
      }
    />
    <Route
      path='/docs/rotate'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Rotate</PageTitle>
          <Rotate />
        </>
      }
    />
    <Route
      path='/docs/ribbon'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Ribbon</PageTitle>
          <Ribbon />
        </>
      }
    />
    <Route
      path='/docs/separator'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Separator</PageTitle>
          <Separator />
        </>
      }
    />
    <Route
      path='/docs/symbol'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Symbol</PageTitle>
          <Symbol />
        </>
      }
    />
    <Route
      path='/docs/tables'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Tables</PageTitle>
          <Tables />
        </>
      }
    />
    <Route
      path='/docs/tabs'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Tabs</PageTitle>
          <Tabs />
        </>
      }
    />
    <Route
      path='/docs/utilities'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Utilities</PageTitle>
          <Utilities />
        </>
      }
    />
    <Route
      path='/docs/nouislider'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>noUISlider</PageTitle>
          <NouiSlider />
        </>
      }
    />

    <Route
      path='/docs/icons/bootstrap'
      element={
        <>
          <PageTitle breadcrumbs={baseBC}>Bootstrap Icons</PageTitle>
          <BootstrapIcons />
        </>
      }
    />

    <Route
      path='/docs/icons/duotune'
      element={
        <>
          <PageTitle breadcrumbs={iconsBC}>Duotune</PageTitle>
          <DuotuneIcons />
        </>
      }
    />
    <Route
      path='/docs/icons/fontawesome'
      element={
        <>
          <PageTitle breadcrumbs={iconsBC}>Font Awesome Icons</PageTitle>
          <FontAwesomeIcons />
        </>
      }
    />
    <Route
      path='/docs/icons/lineawesome'
      element={
        <>
          <PageTitle breadcrumbs={iconsBC}>Line Awesome Icons</PageTitle>
          <LineAwesomeIcons />
        </>
      }
    />
    <Route path='/docs' element={<Navigate to='/docs/quick-start' />} />
    <Route path='/' element={<Navigate to='/docs/quick-start' />} />
  </Routes>
)

export {DocsPage}
