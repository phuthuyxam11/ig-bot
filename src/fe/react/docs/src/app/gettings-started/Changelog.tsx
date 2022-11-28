import {KTSVG} from '../../_metronic/helpers'

export function Changelog() {
  const {REACT_APP_THEME_NAME} = process.env

  return (
    <>
      <div className='accordion accordion-flush accordion-icon-toggle' id='kt_accordion'>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex '
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_1_0'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.1.0 - July 5, 2022
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_1_38'
            className='fs-6 my-1 py-0 ps-10 show collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Added React<code>Demo 6</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo6/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    Added React<code>Demo 7</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo7/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    Added React<code>Demo 8</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo8/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    Runtime switcher between <code>dark</code> and <code>light</code> modes{' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo1/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Updates:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Update <code>React</code> version to <code>v18.0.0</code>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Fixes:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    <code>Yarn</code> and <code>NPM</code> installation issues fixes{' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/docs/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    <code>Demo 1</code> <code>aside toggle</code> behaviour fixes{' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo1/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    <code>Splash screen</code> fixes. Smooth transitions and avoiding blinks till
                    the main application loads{' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo1/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex '
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_38'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.38 - March 30, 2022
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_38'
            className='fs-6 my-1 py-0 ps-10 show collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Added React<code>Demo 4</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo4/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    Added React<code>Demo 5</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo5/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Updates:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Removed <code>Redux</code> boilerplate code for Authorization module.
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex '
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_37'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.37 - March 3, 2022
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_37'
            className='fs-6 my-1 py-0 ps-10 show collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Added React<code>Demo 3</code>
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo3/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>
                    Added <code>User management</code> module (with{' '}
                    {/* <Link to='/docs/crud'>CRUD + Datatable</Link> examples) */}
                    CRUD + Datatable examples){' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo1/apps/user-management/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Updates:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Sync <code>SASS</code> files from the latest <code>HTML</code> version
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_36'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.36 - February 12, 2022
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_36'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Progress bar on route change.{' '}
                    <a
                      href='https://preview.keenthemes.com/metronic8/react/demo1/'
                      className={'ml-2'}
                    >
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Updates:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>React router update.</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_35'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.35 - January 11, 2022
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_35'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    New right toolbar buttons.{' '}
                    <a
                      className='mb-2'
                      href='https://preview.keenthemes.com/metronic8/react/demo1/'
                    >
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Updates:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Demo2 chat indicator position fixes.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo2/'>Preview</a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Fixes:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Authentication bug fixes.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1/auth/registration'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_34'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.34 - December 31, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_34'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Update:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Updated accordion doc.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/docs/docs/accordion'>
                      Preview
                    </a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Fix:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>Fixed undefined access_token.</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_32'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.32 - December 15, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_32'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Added password meter example on Sign up page.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1/auth/registration'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_31'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.31 - November 25, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_31'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Update:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Authentication system improvements.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1'>Preview</a>
                  </li>
                </ul>
                <h3 className='fs-6 fw-bolder mb-1'>Fix:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Layout builder fixes.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1/builder'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_30'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.30 - November 11, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_30'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Fix:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Fixed demo2 aside menu on mobile.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo2'>Preview</a>
                  </li>
                  <li className='py-2'>
                    Fixed create app modal error messages.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1'>Preview</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_29'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.29 - October 29, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_29'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Fix:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Fixed Sign Out button in user menu.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1'>Preview</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_24'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.28 - October 5, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_24'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    New search menu in header.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1'>Preview</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_24'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.25 - September 20, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_24'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Update:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Updated Explore drawer.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1'>Preview</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_24'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.24 - September 2, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_24'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Update:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>SwapperComponent improvements.</li>
                </ul>
              </div>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Fixes:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Fixed breadcrumbs on mobile version.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/demo1/dashboard'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_23'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.23 - August 20, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_23'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    RTL version setup doc.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/docs/docs/rtl'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Update:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Migration to new doutone icons set.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/docs/docs/icons/duotune'>
                      Preview
                    </a>
                  </li>
                  <li className='py-2'>Migration to new illustrations set.</li>
                  <li className='py-2'>DrawerComponent improvements.</li>
                </ul>
              </div>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Fixes:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Duplicated drawer overlay on second DrawerComponent initialization.
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_22'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.22 - August 6, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_22'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Added demo2 react version.{' '}
                    <a href='https://preview.keenthemes.com/metronic/react/demo1/dashboard'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v8_0_21'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.21 - July 30, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v8_0_21'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>New:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    Dark mode setup doc.{' '}
                    <a href='https://preview.keenthemes.com/metronic8/react/docs/docs/dark-mode'>
                      Preview
                    </a>
                  </li>
                </ul>
              </div>
              <div className='mb-5'>
                <h3 className='fs-6 fw-bolder mb-1'>Fix:</h3>
                <ul className='my-0 py-0'>
                  <li className='py-2'>
                    <code className='ms-0'>Error page </code> redirect issue. Now we have page
                    refresh after redirect from Error pages.
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <div className='accordion-item mb-5'>
          <div
            className='accordion-header py-3 d-flex collapsed'
            data-bs-toggle='collapse'
            data-bs-target='#kt_accordion_body_v1_0_0'
            aria-expanded='true'
          >
            <span className='accordion-icon'>
              <KTSVG
                path='/media/icons/duotune/arrows/arr064.svg'
                className='svg-icon svg-icon-3'
              />
            </span>
            <h3 className='fs-3 text-gray-800 fw-bolder mb-0 ms-4'>
              {REACT_APP_THEME_NAME} v8.0.19 - July 20, 2021
            </h3>
          </div>
          <div
            id='kt_accordion_body_v1_0_0'
            className='fs-6 my-1 py-0 ps-10 collapse'
            data-bs-parent='#kt_accordion'
          >
            <div className='accordion-body ps-0 pt-0'>
              <ul className='my-0 py-0'>
                <li className='py-2'>Initial release</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
