## Test Case: Job posting (free plan)
- **Test case ID:** TCP-JOB-MGMT-001
- **Title:** Verify a successful job post for a free plan account
- **Related requirements:** JOB-MGMT-001, JOB-MGMT-001a, JOB-MGMT-001c
- **Related user stories:** US-JOB-MGMT-001
- **Preconditions**
	- The company has logged in to its free plan account.
	- The company has posted 2 of the 3 available job offers.
	- The user has access to the job creation form.
- **Process steps:**
	- On the side navigation bar, click on the "Empleos" option.
	- Click on the "Nuevo trabajo" button above the table displayed.
	- Fill in all the mandatory fields of the job creation form.
	- Click on the "Publicar" button
- **Expected result**
	- The system processes the request without errors.
	- The new job appears in the jobs table immediately after publishing.
	- The system displays a success message.

## Test Case: Job Posting Limit (free plan)
- **Test Case ID:** TCP-JOB-MGMT-002
- **Title:** Verify the system prevents posting more than 3 offers with the free plan
- **Related requirements:** REQ-JOB-MGMT-001, REQ-JOB-MGMT-001a, REQ-JOB-MGMT-001c
- **Related user stories:** US-JOB-MGMT-001
- **Preconditions**
	- The user has logged in to its free plan account
	- The user has already posted 3 of 3 available job offers.
- **Process steps**
	- Navigate to "Empleos" page
	- Move the mouse over the "nuevo trabajo" button
- **Expected result**
	- The "nuevo trabajo" button is disabled (grayed out).
	- A pop-up message appears on hover stating that a premium account is required to add more jobs

## Test Case: Job Posting(premium plan)
- **Test Case ID:** TCP-JOB-MGMT-003
- **Title:** Verify successful job posting for a premium plan account
- **Related requirements:** REQ-JOB-MGMT-001, REQ-JOB-MGMT-001b, REQ-JOB-MGMT-001c
- **Related user stories:** US-JOB-MGMT-001
- **Precondition**
	- The company has already logged in to its premium plan account.
	- The company has posted at least 3 job offers.
- **Process steps**
	- Navigate to "Empleos" page
	- Click on "nuevo trabajo" button
	- Fill in all the mandatory fields of the job creation form
	- Click on the "Publicar" button
- **Expected result**
	- The system processes the request without errors.
	- The new job is displayed in the table.
	- A success message is displayed.

## Test Case: Mandatory field validation
- **Test Case ID:** TCP-JOB-MGMT-004
- **Title:** Verify that the system validates mandatory fields before submitting the job post.
- **Related requirements:** REQ-JOB-MGMT-001, REQ-JOB-MGMT-001c
- **Related user stories:** US-JOB-MGMT-001
- **Precondition**
	- The company has already logged in to its account (free or premium)
	- The user is on the "Empleos" page
- **Process steps**
- Click on "nuevo trabajo" button
- Leave one of the mandatory fields empty (e.g., "Title").
- Fill in all the remaining mandatory fields
- Click the "publicar" button
- **Expected result**
	- The system does not allow the form to be submitted.
	- The system mark on red, the empty mandatory field
