# Auth requirements

### REQ-AUTH-001 - The user can create an account

- **Description** The system must enable a user to create a new account

- **Priority** High

- **Acceptance criteria**
	- The system displays a form with all the required fields for a new user registration
	- The system successfully saves the user after clicking on 'Registrar'
	- The system redirects the user to the login 
	- The system shows a message confirming the user registration
- **Related requirements**
	- **REQ-AUTH-001a - user type account**: the registration form registers a new user as a "job seeker"
		- **Acceptance criteria**
			- The system prevents the registration of the new user if the email already exists
	- **REQ-AUTH-001b - New user fields:** The new user registration form must request the following fields: Name, Last name, Email, Phone, Password, and Confirm Password
		- **Acceptance criteria**
			- The registration form contains all the required fields
			- The system validates if all the mandatory fields are filled before the registration