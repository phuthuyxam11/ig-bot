/* eslint-disable react/jsx-no-target-blank */
/* eslint-disable jsx-a11y/anchor-has-content */
import React from 'react'

export function References() {
  const {REACT_APP_THEME_NAME} = process.env

  return (
    <>
      <h1 className='anchor fw-bolder mb-2' id='overview'>
        <a href='#overview'></a>Overview
      </h1>

      <div className='pt-5'>
        <strong>{REACT_APP_THEME_NAME}</strong>&nbsp; uses the following open source resources:
      </div>
      <div className='py-5'>
        <div className='table-responsive border rounded'>
          <table className='table table-striped table-flush align-middle mb-0'>
            <thead>
              <tr className='fs-4 fw-bolder text-dark p-6'>
                <th className='min-w-400px'>Name</th>
                <th className='min-w-400px'>Link</th>
                <th className='min-w-100px'>Version</th>
              </tr>
            </thead>
            <tbody>
              <tr className='p-6'>
                <td>
                  <code>Node.js (LTS version)</code>
                </td>
                <td>
                  <a href='https://nodejs.org/en/' target='_blank'>
                    https://nodejs.org/en/
                  </a>
                </td>
                <td>v16.15.1</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React.js</code>
                </td>
                <td>
                  <a href='https://en.reactjs.org/' target='_blank'>
                    https://en.reactjs.org/
                  </a>
                </td>
                <td>^18.0.0</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React Router DOM</code>
                </td>
                <td>
                  <a href='https://reactrouter.com/' target='_blank'>
                    https://reactrouter.com/
                  </a>
                </td>
                <td>^6.3.0</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React query</code>
                </td>
                <td>
                  <a href='https://react-query.tanstack.com/' target='_blank'>
                    https://react-query.tanstack.com/
                  </a>
                </td>
                <td>^3.38.0</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React Table</code>
                </td>
                <td>
                  <a href='https://tanstack.com/table/v8/docs/examples/react/basic' target='_blank'>
                    https://tanstack.com/table/v8/docs/examples/react/basic
                  </a>
                </td>
                <td>^7.7.0</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>TypeScript</code>
                </td>
                <td>
                  <a href='https://www.typescriptlang.org/' target='_blank'>
                    https://www.typescriptlang.org/
                  </a>
                </td>
                <td>4.6.3</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>Axios</code>
                </td>
                <td>
                  <a href='https://github.com/axios/axios' target='_blank'>
                    https://github.com/axios/axios
                  </a>
                </td>
                <td>^0.26.1</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>Bootstrap</code>
                </td>
                <td>
                  <a href='https://getbootstrap.com/docs/5.0' target='_blank'>
                    https://getbootstrap.com/docs/5.0
                  </a>
                </td>
                <td>5.2.0-beta1</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React Bootstrap v5</code>
                </td>
                <td>
                  <a href='https://react-bootstrap-v5.netlify.app/' target='_blank'>
                    https://react-bootstrap-v5.netlify.app/
                  </a>
                </td>
                <td>2.3.0</td>
              </tr>

              <tr className='p-6'>
                <td>
                  <code>Formik</code>
                </td>
                <td>
                  <a href='https://formik.org/' target='_blank'>
                    https://formik.org/
                  </a>
                </td>
                <td>^2.2.9</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>React ApexChart</code>
                </td>
                <td>
                  <a href='https://apexcharts.com/docs/react-charts/' target='_blank'>
                    https://apexcharts.com/docs/react-charts/
                  </a>
                </td>
                <td>^1.3.9</td>
              </tr>
              <tr className='p-6'>
                <td>
                  <code>Chart JS</code>
                </td>
                <td>
                  <a href='https://www.chartjs.org/' target='_blank'>
                    https://www.chartjs.org/
                  </a>
                </td>
                <td>^3.3.2</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </>
  )
}
