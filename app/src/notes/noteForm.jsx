import React from 'react'

export default (props) => {
    return (
            <form className="mb-5 ml-4 col-md-6" onSubmit={e => e.preventDefault()}>
                
                <div className="form-group">
                    <label htmlFor="headline">Title</label>
                    <input type="text" name="headline" 
                        className="form-control" 
                        value={props.headline}
                        onChange={props.onChange} />
                </div>

                 <div className="form-group">
                    <label htmlFor="content">Description</label>
                    <textarea name="content" 
                        value={props.content}
                        className="form-control" 
                        onChange={props.onChange} ></textarea>
                </div>
                <div className="form-group">
                <div className="form-group">
                        <label htmlFor="reminder">Remind</label>
                        <input type="datetime" name="reminder" 
                            className="form-control" 
                            value={props.reminder}
                            onChange={props.onChange} />
                </div>

                <div className="row d-flex align-items-center">
                    <div className="col-md-12 float-right">
                
                       <button className="btn btn-primary" onClick={props.onClick}>Confirmar</button>
                    </div>
                    </div>
                </div>
            </form>
    )
} 