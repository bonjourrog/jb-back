# Job Management requirements

### REQ-JOB-MGMT-001 - The company can post a job

- **Description**: The system must enable a company to create and publish a new job posting.

- **Priority:** High

- **Acceptance criteria**
	- The user is logged in to his account (free or premium)
	- The system displays a form with all the required fields for a new job posting.
	- The system successfully saves the job posting after the user clicks "Publicar".
	- The system displays a confirmation message to the user after a successful post
- **Related requirements**
	- **REQ-JOB-MGMT-001a - free plan:** Companies on a free plan can post a maximum of 3 jobs.
		- **Acceptance criteria**
			- The system validates the number of jobs posted by a company on a free plan
			- The system prevents the creation of a new job if the limit of 3 jobs is reached.
			- The system disable "Nuevo trabajo" button (gray-out) and pop ups a message indicating that is necesarry a premuim account
	- **REQ-JOB-MGMT-001b - Premium plan**: Companies on a premium plan can post an unlimited number of jobs.
		- **Acceptance criteria**
			- The system allows companies on a premium plan to post any number of jobs without restrictions
	- **REQ-JOB-MGMT-001c - job post fields:** Each job post must include the following fields: title, description, short description, salary, benefits, location, industry, schedule, and contract type
		- **Acceptance criteria**
			- The job creation form contains all the required fields
			- The system validates that all mandatory fields are filled before saving the post

### RQE-JOB-MGMT-002 - The company can delete a job

- **Description:** The system must allow a company to delete their own job posting.

- **Priority**: High

- **Acceptance criteria**
	- The user is logged in and owns a at least 1 job
	- The system displays a "Delete" option for the company's jobs
	- The system prompts for confirmation before deletion
	- The job posting is permanently removed from the list after confirmation
	- The job posting is no longer visible to job seekers after deletion
	- The system displays a confirmation message after successful deletion

### RQE-JOB-MGMT-003 - The company can update job information

- **Description:** The system must allow a company to edit and update their existing job postings.

- **Priority:** High

- **Acceptance criteria**
- The user is logged in and owns a at least 1 job
- The system displays an "Edit" option for the company's job postings
- The system shows a pre-populated form with current job information
- The system validates all required fields before saving changes
- The updated information is immediately visible to job seekers
- The system records the last modification date
- The system displays a confirmation message after successful update

### RQE-JOB-MGMT-004 - The company can disable the job

- **Description:** The system must allow a company to temporarily disable a job posting without deleting it.

- **Priority:** Medium

- **Acceptance criteria**
- The user is logged in and owns a at least 1 job
- The system displays a "Publicado" toggle for job postings
- Disabled jobs are hidden from job seeker searches and listings
- Disabled jobs remain in the company's job management panel with "Inactive" status
- The company can reactivate disabled jobs at any time
- The system maintains all application data for disabled jobs
- The system displays the current status (Active/Inactive) clearly

### RQE-JOB-MGMT-005 - The company can update the password
- **Description:** The system must allow companies to change their account password with proper security validation.

- **Priority:** Medium

- **Acceptance criteria**
- The user is logged in to their company account
- The system displays a "Change Password" form in account settings
- The system requires the current password for verification
- The new password must meet security requirements (minimum length, complexity)
- The system confirms password change with a double-entry field
- All active sessions are logged out after password change (optional security measure)
- The system displays a success message after password update
- The system sends email notification about password change

### RQE-JOB-MGMT-006 - The company can list all jobs created

- **Description:** The system must provide a comprehensive view of all job postings created by the company.

- **Priority:** High

- **Acceptance criteria**
- The user is logged in to their company account
- The system displays a paginated list of all company's job postings
- Each job entry shows: title, status (Active/Inactive), creation date, and number of applications
- The system provides filtering options (by status, date range, industry)
- The system provides sorting options (by date, title, application count)
- The list shows both active and inactive job postings with clear status indicators
- Pagination controls are displayed when there are more than 20 jobs
- The system displays total count of jobs

### RQE-JOB-MGMT-007 - The company can list all user that applied to each job

- **Description:** The system must allow companies to view all candidates who have applied to their job postings.

- **Priority:** High

- **Acceptance criteria**
- The user is logged in and owns the job posting
- The system displays an "Applications" or "Candidates" section for each job
- The list shows applicant name, application date, and current status
- The system provides filtering options by application status (new, under review, rejected, hired)
- The system displays the total count of applications
- Applications are sorted by date (newest first) by default
- The system provides pagination for jobs with many applications
- Each application entry has a link to view full candidate details

### RQE-JOB-MGMT-008 - The company can check the information of users who applied

- **Description:** The system must allow companies to view detailed information about candidates who applied to their jobs.

- **Priority:** High

- **Acceptance criteria**
- The user is logged in and owns the job posting the candidate applied to
- The system displays a detailed candidate profile when selected from applications list
- The profile includes: resume/CV, contact information, cover letter, and application responses
- The system respects candidate privacy settings and only shows authorized information
- The system displays application date and any previous communication history
- The profile shows the candidate's application status for this specific job
- The system provides options to download candidate's resume (if permitted)
- The system logs when company views candidate information for audit purposes


### RQE-JOB-MGMT-009 - The company can change the job status (in review, rejected, etc)

- **Description:** The system must allow companies to update the application status for candidates who applied to their jobs.

- **Priority:** High

- **Acceptance criteria**
- The user is logged in and owns the job posting
- The system provides a status dropdown/selector for each application
- Available statuses include: New, Under Review, Pre-selected, Interview Scheduled, Rejected, Hired
- Status changes are immediately saved and reflected in the application list
- The system records status change history with timestamp and user who made the change
- The system optionally sends notification to candidate when status changes (configurable)
- Bulk status update option is available for multiple applications
- The system prevents invalid status transitions (business rules apply)

### RQE-JOB-MGMT-010 - The company can contact with a user who is appliying

- ** Description:** The system must provide communication channels for companies to contact candidates.

- **Priority:** Medium

- **Acceptance criteria**
- The user is logged in and owns the job posting the candidate applied to
- The system displays a "Contact" button/option for each application
- The system provides email communication option through the platform
- The system optionally provides in-app messaging functionalityw
- All communication is logged and visible to both parties
- The system records when contact was initiated and by whom
- Email templates are available for common communications (interview invitation, rejection, etc.)
- The system respects candidate's communication preferences
- Contact history is maintained for each application




