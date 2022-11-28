/* eslint-disable react/jsx-no-target-blank */
/* eslint-disable jsx-a11y/anchor-has-content */
import {FC} from 'react'
import {KTSVG} from '../../_metronic/helpers'

const Backend: FC = () => (
  <>
    <h1 className='anchor fw-bolder mb-2' id='how-to'>
      <a href='#how-to'></a>Back-end
    </h1>
    <div className='pt-5'>
      We've prepared our own example of Authorization and User-management API, you can easily to
      switch for your own in the <code>demo1/.env</code> file.
    </div>
    <div className='pt-5'></div>
    <div className='d-flex align-items-center rounded py-5 px-5 bg-light-danger'>
      <KTSVG
        path='/media/icons/duotune/general/gen044.svg'
        className='svg-icon-3x svg-icon-danger me-5'
      />
      <div className='text-gray-700 fw-bold fs-6'>
        Info: For User-management module, you can use generated <a href='https://preview.keenthemes.com/theme-api/api/documentation'>Swagger</a> documentation.
      </div>
    </div>
  </>
)

export {Backend}
