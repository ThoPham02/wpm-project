import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

function FormRegister() {
  return (
    <Form className="formLogin">
      <Form.Group className="mb-3" controlId="form-name">
        <Form.Label>User Name</Form.Label>
        <Form.Control type="email" placeholder="Enter email" />
      </Form.Group>

      <Form.Group className="mb-3" controlId="form-email">
        <Form.Label>Email</Form.Label>
        <Form.Control type="password" placeholder="Mail" />
      </Form.Group>

      <Form.Group className="mb-3" controlId="form-password">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" />
      </Form.Group>

      <Form.Group className="mb-3" controlId="form-passwordConfirm">
        <Form.Label>Password Confirm</Form.Label>
        <Form.Control type="password" placeholder="Password confirm" />
      </Form.Group>

      <Button variant="primary" type="submit" className="formLogin-btn">
        Submit
      </Button>
    </Form>
  );
}

export default FormRegister;
