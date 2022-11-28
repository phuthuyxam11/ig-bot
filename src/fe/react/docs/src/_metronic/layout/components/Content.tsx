import {FC} from 'react'
import {WithChildren} from '../../helpers'

const Content: FC<WithChildren> = ({children}) => (
  <div className='container' id='kt_docs_content_container'>
    <div className='card mb-2'>
      <div className='card-body fs-6 py-15 px-10 py-lg-15 px-lg-15'>{children}</div>
    </div>
  </div>
)

export {Content}
