# Frontend-User Implementation Checklist

## 📋 Quick Reference

### Endpoints Required from Backend
- `POST /auth/login/rumah-sakit` → Login
- `GET /rumah-sakit/me` → Get hospital profile
- `PUT /rumah-sakit/me` → Update hospital profile
- `GET /permintaan-darah/my-requests` → List my requests
- `POST /permintaan-darah` → Create request
- `GET /permintaan-darah/:id` → Get request detail
- `PUT /permintaan-darah/my-requests/:id` → Update own request
- `DELETE /permintaan-darah/:id` → Delete request
- `PUT /permintaan-darah/update/:id` → Update request status
- `GET /dashboard/status-summary/{hospitalId}` → Dashboard stats
- `GET /ws/connect` → WebSocket connection

---

## 🔧 SETUP PHASE

### Project Setup
- [ ] Clone/create frontend-user project
- [ ] Install dependencies (npm install)
- [ ] Setup TypeScript config
- [ ] Setup Vite config
- [ ] Setup Tailwind CSS
- [ ] Copy design tokens from frontend-admin

### Development Environment
- [ ] .env.local setup (API_BASE_URL, etc)
- [ ] ESLint & Prettier config
- [ ] Git setup (.gitignore, etc)
- [ ] VS Code extensions (Volar, Tailwind, etc)

### Frontend-Admin Code Reuse
- [ ] Copy axios client setup
- [ ] Copy Pinia store structure
- [ ] Copy layout components
- [ ] Copy UI components
- [ ] Copy utility functions
- [ ] Copy type definitions

---

## 🔐 AUTHENTICATION

### Login Implementation
- [ ] Create LoginView.vue
- [ ] Form validation (client-side)
- [ ] API call to login endpoint
- [ ] Token storage (localStorage or cookie)
- [ ] Store user info in auth store
- [ ] Redirect to dashboard on success
- [ ] Show error message on failure
- [ ] Handle rate limiting (429 response)
- [ ] Remember me functionality (optional)

### Route Protection
- [ ] Create router guards
- [ ] Redirect unauthenticated to login
- [ ] Add Authorization header to API calls
- [ ] Handle 401 → redirect to login
- [ ] Handle 403 → show access denied
- [ ] Auto-logout on token expiry

### Logout
- [ ] Clear token from storage
- [ ] Clear user from store
- [ ] Clear all stores/caches
- [ ] Redirect to login

---

## 📊 DASHBOARD

### Dashboard View
- [ ] Create DashboardView.vue
- [ ] Fetch dashboard data (status summary)
- [ ] Display 4 stat cards (total, pending, completed, cancelled)
- [ ] Add status distribution pie chart
- [ ] Add requests over time line chart
- [ ] Show recent requests list
- [ ] Add quick action buttons
- [ ] Show WebSocket connection status
- [ ] Responsive layout (mobile/tablet/desktop)

### Real-time Updates
- [ ] Show connection indicator
- [ ] Update stats on new request
- [ ] Update stats on status change
- [ ] Show toast notification on update

### Charts & Visualization
- [ ] Setup chart library (Chart.js or similar)
- [ ] Pie chart for status distribution
- [ ] Line chart for requests over time
- [ ] Responsive chart sizing
- [ ] Loading state for charts

---

## 📋 MY REQUESTS - LIST VIEW

### Requests List Page
- [ ] Create RequestsListView.vue
- [ ] Fetch requests with filtering
- [ ] Display table with columns (ID, Patient, Blood Type, Quantity, Status, Date, Actions)
- [ ] Implement sorting (by date, status, patient name)
- [ ] Implement filtering:
  - [ ] By status (dropdown)
  - [ ] By date range (date pickers)
  - [ ] Search by patient name or request ID

### Table Features
- [ ] Pagination (20 per page)
- [ ] Status badge colors
- [ ] Action buttons (View, Edit, Cancel)
- [ ] Loading skeleton
- [ ] Empty state message
- [ ] Error message display

### Search & Filter
- [ ] Status filter dropdown
- [ ] Date range picker
- [ ] Search input
- [ ] Apply filters button
- [ ] Reset filters button
- [ ] Show active filters count

---

## 🔍 REQUEST DETAIL

### Detail View
- [ ] Create RequestDetailView.vue
- [ ] Fetch request data
- [ ] Display patient information
- [ ] Display blood request details
- [ ] Display requested components list
- [ ] Display current status (prominent)

### Status Timeline
- [ ] Create RequestStatusTimeline component
- [ ] Display all status changes
- [ ] Show timestamps for each change
- [ ] Show who made the change (admin name)
- [ ] Show reason for cancellation (if applicable)
- [ ] Visual timeline design

### Actions
- [ ] View button (always available)
- [ ] Edit button (if status allows)
- [ ] Mark as Picked Up button (if bisa_diambil)
- [ ] Cancel button with reason input (if cancellable)
- [ ] Back to list button
- [ ] Confirmation dialogs for actions

---

## ➕ CREATE/EDIT REQUEST

### Create Request Form
- [ ] Create CreateEditRequestView.vue
- [ ] Section 1: Patient Information
  - [ ] Patient name (required)
  - [ ] Medical record number (required)
  - [ ] Birthplace
  - [ ] Birth date (required)
  - [ ] Gender (required)
  - [ ] Contact number

- [ ] Section 2: Blood & Transfusion Info
  - [ ] Blood type (required)
  - [ ] Rhesus factor (required)
  - [ ] Hemoglobin level (optional)
  - [ ] Indication for transfusion
  - [ ] Previous transfusion (yes/no)
  - [ ] Pregnancy history (if applicable)

- [ ] Section 3: Request Components
  - [ ] Component dropdown (from backend)
  - [ ] Blood type (auto-filled)
  - [ ] Quantity (required)
  - [ ] Add/remove component rows

- [ ] Section 4: Medical Info
  - [ ] Ward/department
  - [ ] Room/bed number
  - [ ] Doctor name
  - [ ] Doctor contact
  - [ ] Additional notes

### Form Features
- [ ] Real-time validation
- [ ] Error messages (inline)
- [ ] Blood type auto-propagation
- [ ] Dynamic component addition
- [ ] Form reset button
- [ ] Cancel button
- [ ] Submit button with loading state
- [ ] Success message
- [ ] Pre-fill from hospital profile (for edit)

### Edit Functionality
- [ ] Fetch existing request
- [ ] Pre-populate form fields
- [ ] Disable fields if request not editable
- [ ] Show warning if editing processed request
- [ ] Update API call on submit
- [ ] Success/error handling

---

## 👤 PROFILE MANAGEMENT

### Profile View
- [ ] Create ProfileView.vue
- [ ] Fetch hospital profile (GET /rumah-sakit/me)
- [ ] Display hospital information (readonly)
- [ ] Display account information
- [ ] Show last login date

### Edit Profile
- [ ] Add Edit button
- [ ] Enable edit mode for fields
- [ ] Hospital name (text)
- [ ] Address (textarea)
- [ ] Phone number
- [ ] Email
- [ ] Logo/image upload (optional)
- [ ] Save button with loading state
- [ ] Cancel button (discard changes)
- [ ] Success/error messages

### Change Password
- [ ] Create password change section
- [ ] Current password (required)
- [ ] New password (required)
- [ ] Confirm password (required)
- [ ] Password strength indicator
- [ ] Show password requirements
- [ ] Change password button
- [ ] Success/error messages

### Validation
- [ ] Required fields
- [ ] Email format
- [ ] Phone number format
- [ ] Password strength (min 8 chars, etc)
- [ ] Confirm password match

---

## 🔄 REAL-TIME FEATURES

### WebSocket Connection
- [ ] Setup WebSocket service
- [ ] Connect on app load
- [ ] Endpoint: `/ws/connect?token=JWT_TOKEN`
- [ ] Handle connection success
- [ ] Handle connection error
- [ ] Auto-reconnect on disconnect
- [ ] Heartbeat/ping-pong

### Message Handling
- [ ] Listen for "request_status_changed" messages
- [ ] Listen for "notification" messages
- [ ] Update local store on message
- [ ] Refresh relevant data
- [ ] Show toast notifications

### Connection Indicator
- [ ] Show "Online" (green dot) when connected
- [ ] Show "Offline" (gray dot) when disconnected
- [ ] Place in header
- [ ] Show reconnection attempts

---

## 🎨 UI COMPONENTS

### Reused from frontend-admin
- [ ] AppLayout component
- [ ] AppHeader component
- [ ] AppSidebar component
- [ ] Button component (all variants)
- [ ] TextInput component
- [ ] SelectInput component
- [ ] DateInput component
- [ ] RadioInput component
- [ ] TextArea component
- [ ] Table component
- [ ] Card component
- [ ] Badge component
- [ ] Modal component
- [ ] Toast/Notification component
- [ ] Loading skeleton
- [ ] Empty state

### New Components
- [ ] RequestStatusTimeline
- [ ] BloodTypeSelector
- [ ] ComponentRequestForm
- [ ] RequestCard
- [ ] HospitalBranding

---

## 🧪 TESTING

### Unit Tests
- [ ] Store actions (auth, my-requests, my-profile, dashboard)
- [ ] Utility functions (date, validation, etc)
- [ ] Component logic (computed, methods)

### Integration Tests
- [ ] API mocking (mock axios)
- [ ] Store updates after API calls
- [ ] Route navigation
- [ ] Auth flow

### E2E Tests
- [ ] Login flow
- [ ] Create request flow
- [ ] View requests list
- [ ] Edit request
- [ ] Cancel request
- [ ] View profile
- [ ] Edit profile
- [ ] Change password
- [ ] Logout

---

## 📱 RESPONSIVE DESIGN

### Mobile (< 640px)
- [ ] Stack layout vertically
- [ ] Full-width forms
- [ ] Hamburger menu for sidebar
- [ ] Touch-friendly buttons (44x44px min)
- [ ] Readable fonts (16px min)
- [ ] Horizontal scroll tables
- [ ] Bottom navigation (optional)

### Tablet (640px - 1024px)
- [ ] 2-column layouts where applicable
- [ ] Sidebar collapsible
- [ ] Charts responsive
- [ ] Optimized spacing

### Desktop (> 1024px)
- [ ] Full sidebar
- [ ] Multi-column layouts
- [ ] Full-featured charts
- [ ] Proper spacing

---

## ♿ ACCESSIBILITY

### Standards Compliance
- [ ] WCAG 2.1 Level AA
- [ ] Semantic HTML
- [ ] Keyboard navigation (Tab, Enter, Esc)
- [ ] ARIA labels & roles
- [ ] Color contrast (4.5:1)
- [ ] Focus indicators
- [ ] Form labels associated

### Implementation
- [ ] Use semantic tags (button, form, input, label)
- [ ] Add aria-labels for complex elements
- [ ] Skip to main content link
- [ ] Test with screen reader
- [ ] High contrast colors
- [ ] Focus visible on all interactive elements

---

## 🚀 PERFORMANCE

### Optimization
- [ ] Lazy load routes
- [ ] Code splitting for pages
- [ ] Image optimization
- [ ] Minimize bundle size
- [ ] Tree-shake unused code
- [ ] Debounce search/filter inputs
- [ ] Cache API responses appropriately
- [ ] Minify CSS/JS for production

### Monitoring
- [ ] Setup error tracking
- [ ] Monitor API performance
- [ ] Track page load times
- [ ] Monitor WebSocket health

---

## 🔒 SECURITY

### Implementation
- [ ] Store token securely (httpOnly cookie preferred)
- [ ] Validate all user inputs
- [ ] Sanitize user input before display
- [ ] Use HTTPS only
- [ ] Implement CSRF protection
- [ ] Don't expose sensitive data in logs
- [ ] Proper error handling (don't expose system details)

### Data Protection
- [ ] Encrypted password transmission
- [ ] Secure session management
- [ ] Rate limiting (inherited from backend)
- [ ] Input validation & sanitization

---

## 📦 DEPLOYMENT

### Build Process
- [ ] Build command: `npm run build`
- [ ] Build output: `/dist`
- [ ] Environment variables in production build
- [ ] Source maps handling

### Deployment Checklist
- [ ] Environment variables configured
- [ ] API endpoints correct for environment
- [ ] Token storage secure
- [ ] Error tracking enabled
- [ ] Performance monitoring enabled
- [ ] Security headers configured
- [ ] CORS properly configured

### Post-Deployment
- [ ] Smoke testing on production
- [ ] Monitor for errors
- [ ] Check performance metrics
- [ ] Verify WebSocket connection
- [ ] Test on various browsers/devices

---

## ✅ FINAL CHECKLIST

### Code Quality
- [ ] ESLint passes
- [ ] No console warnings/errors
- [ ] TypeScript strict mode passes
- [ ] Code formatted with Prettier
- [ ] Meaningful git commits

### Testing
- [ ] All unit tests pass
- [ ] All integration tests pass
- [ ] All E2E tests pass
- [ ] Manual testing completed
- [ ] Cross-browser testing done
- [ ] Mobile testing done

### Documentation
- [ ] README with setup instructions
- [ ] Component documentation
- [ ] API integration docs
- [ ] Environment variables documented
- [ ] Deployment guide

### Performance & Security
- [ ] Lighthouse score > 80
- [ ] No security vulnerabilities (npm audit)
- [ ] Load time < 3 seconds
- [ ] Responsive design working
- [ ] Accessibility audit passed

### Production Ready
- [ ] All issues resolved
- [ ] Performance optimized
- [ ] Security reviewed
- [ ] Documentation complete
- [ ] Ready for deployment

---

**Total Estimated Time**: 14-20 days  
**Created**: 2026-04-28  
**Version**: 1.0
