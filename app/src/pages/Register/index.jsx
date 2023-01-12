import React from 'react';
import { Link } from "react-router-dom";
import FormRegister from './components/FormRegister';

function Register() {
    return (
        <div className='login'>
            <h2 className='login-title'>Register</h2>
            <FormRegister />
            <Link to="/login">Back to login</Link>
        </div>
    );
}

export default Register;