# Frontend-User Design Guidelines & Implementation Roadmap
# Aplikasi Permintaan Darah PMI - Rumah Sakit (RS) User Portal

## 📋 OVERVIEW

Frontend-user adalah portal khusus untuk Rumah Sakit (RS) users yang ingin:
- Membuat permintaan darah baru
- Mengelola permintaan darah mereka
- Melihat dashboard/statistik rumah sakit
- Mengelola profil rumah sakit

**Target User**: Hospital Staff (Rumah Sakit)  
**Tech Stack**: Vue 3 + TypeScript + Tailwind CSS (reuse dari frontend-admin)  
**Auth Role**: `rumah_sakit` (role dari backend)

---

## 🎨 DESIGN SYSTEM & LAYOUT

### Layout Structure
✅ **Header Component**
   - PMI logo/branding
   - Hospital name (from profile)
   - Notifications icon (real-time updates)
   - Settings/user menu
   - Logout button

✅ **Sidebar Navigation** (3-4 items)
   - Dashboard
   - My Requests (main feature)
   - Create Request (quick action)
   - My Profile
   - Logout

✅ **Main Content Area**
   - Dynamic content based on current page
   - Consistent padding with frontend-admin

✅ **Footer** (Optional)
   - Copyright, terms, privacy links

### Design Language
✅ **Reuse from frontend-admin**:
   - Color scheme (blue primary, green success, red danger)
   - Typography (same fonts, sizes)
   - Spacing system
   - Border radius & shadows
   - Icon library (Lucide Vue)

### Status Indicator Colors
```
dibuat (created)           → Gray
diproses (processing)      → Blue
bisa_diambil (ready)       → Green
selesai (completed)        → Dark Green
dibatalkan (cancelled)     → Red
```

---

## 📱 PAGES & MODULES (3-4 Main Pages)

### PAGE 1: LOGIN
**Purpose**: Authenticate RS users  
**Endpoint**: `POST /auth/login/rumah-sakit`

```
Components:
├─ PMI logo
├─ Form (username, password)
├─ "Login" button
├─ "Remember me" checkbox (optional)
├─ Error message display
└─ Loading state

Behavior:
├─ Validate required fields
├─ Show loading spinner
├─ On success: Store token → Redirect to dashboard
├─ On error: Show error message
└─ Rate limiting: Show 429 message if too many attempts
```

---

### PAGE 2: DASHBOARD
**Purpose**: Show RS overview & quick statistics  
**Endpoint**: `GET /dashboard/status-summary/{hospitalId}`

```
Widgets/Cards (4 columns on desktop):
├─ Total Requests (count card)
├─ Pending Requests (count card)
├─ Completed Requests (count card)
└─ Cancelled Requests (count card)

Charts:
├─ Status Distribution (pie chart)
└─ Requests Over Time (line chart - last 30 days)

Recent Section:
├─ Recent Requests (mini table, last 5-10)
└─ Quick Actions (Create New button, View All link)

Real-time:
├─ WebSocket connection indicator
├─ Live status updates
└─ Notification badge

Responsive:
├─ Mobile: Stack cards vertically
├─ Tablet: 2 columns
└─ Desktop: 4 columns
```

---

### PAGE 3: MY REQUESTS (List View)
**Purpose**: List all requests created by this RS  
**Endpoint**: `GET /permintaan-darah/my-requests`

```
Table Features:
├─ Columns: ID | Patient | Blood Type | Quantity | Status | Date | Actions
├─ Pagination (20 per page)
├─ Sorting (date, status, patient name)
├─ Filtering:
│  ├─ By Status (all, dibuat, diproses, bisa_diambil, selesai, dibatalkan)
│  ├─ By Date Range
│  └─ Search (Patient Name or Request ID)
│
├─ Status Color Coding (badges)
├─ Action Buttons:
│  ├─ View (arrow) → detail page
│  ├─ Edit (pencil) → edit form (if editable)
│  └─ Cancel (X) → cancel request (if cancellable)
│
├─ Empty State: "No requests yet" + Create button
└─ Loading State: Skeleton while fetching
```

---

### PAGE 4: REQUEST DETAIL
**Purpose**: View full details & status history  
**Route**: `/requests/:id`  
**Endpoint**: `GET /permintaan-darah/:id`

```
Main Information (left section):
├─ Request ID & Date
├─ Patient Info (name, RM, birthplace, gender)
├─ Blood Type (ABO + Rhesus)
├─ Requested Components (table)
├─ Medical Info (ward, room, doctor, notes)
└─ Current Status (prominent display)

Status Timeline (right section):
├─ Visual timeline (vertical line)
├─ Status changes with timestamps
├─ Who made the change (admin name)
├─ Reason for change (if applicable)
└─ Interactive points on timeline

Actions (based on status):
├─ View: Always available
├─ Edit: Only if dibuat or diproses
├─ Mark as Picked Up: Only if bisa_diambil
├─ Cancel: Only if dibuat, diproses, or bisa_diambil
│  └─ Show reason input field
└─ Back to List button
```

---

### PAGE 5: CREATE/EDIT REQUEST
**Purpose**: Create new or edit existing request  
**Endpoints**: 
- `POST /permintaan-darah` (create)
- `PUT /permintaan-darah/my-requests/:id` (update)

```
SECTION 1: PATIENT INFORMATION
├─ Patient Name (text) *required
├─ Medical Record Number (text) *required
├─ Birthplace (text)
├─ Birth Date (date picker) *required
├─ Gender (radio: L/P) *required
└─ Contact Number (phone)

SECTION 2: BLOOD & TRANSFUSION INFO
├─ Blood Type (dropdown: A, B, AB, O) *required
├─ Rhesus Factor (radio: +/-) *required
├─ Hemoglobin Level (number, optional)
├─ Indication for Transfusion (textarea)
├─ Previous Transfusion? (yes/no)
├─ Transfusion History (if yes)
├─ Ever Pregnant? (yes/no, for female)
└─ Pregnancy Details (if yes)

SECTION 3: REQUEST COMPONENTS
├─ Add Component button (+)
├─ Component List (dynamic rows):
│  ├─ Component Name (dropdown from backend)
│  ├─ Blood Type (auto-filled)
│  ├─ Quantity (number) *required
│  └─ Remove button
└─ Total quantity display

SECTION 4: MEDICAL INFO
├─ Ward/Department (text)
├─ Room/Bed Number (text)
├─ Doctor Name (text)
├─ Doctor Contact (text)
└─ Additional Notes (textarea)

Form Features:
├─ Real-time validation (show errors inline)
├─ Auto-save drafts (optional)
├─ Clear form button
├─ Cancel button (back to list)
├─ Submit button (Create/Update)
├─ Loading state on submit
├─ Success message on submit
└─ Blood type auto-propagates to components
```

---

### PAGE 6: PROFILE
**Purpose**: View and edit RS hospital information  
**Endpoints**:
- `GET /rumah-sakit/me` (view)
- `PUT /rumah-sakit/me` (update)

```
SECTION 1: HOSPITAL INFORMATION (readonly until edit)
├─ Hospital Name
├─ Address
├─ Phone Number
├─ Email
└─ Logo/Image

SECTION 2: ACCOUNT INFORMATION
├─ Username
├─ Email
└─ Last Login

SECTION 3: EDIT FUNCTIONALITY
├─ Edit button (pencil icon)
├─ In edit mode:
│  ├─ Hospital Name (text input)
│  ├─ Address (textarea)
│  ├─ Phone Number (phone input)
│  ├─ Email (email input)
│  └─ Logo/Image upload (optional)
├─ Save button
└─ Cancel button

SECTION 4: SECURITY
├─ Change Password section:
│  ├─ Current Password (password) *required
│  ├─ New Password (password) *required
│  ├─ Confirm Password (password) *required
│  ├─ Change Password button
│  └─ Password requirements (min 8 chars, etc)
└─ Password strength indicator
```

---

## 🔐 AUTHENTICATION & STATE MANAGEMENT

### Authentication Flow
```
1. User logs in (username/password)
2. POST /auth/login/rumah-sakit
3. Backend returns { token, user_info }
4. Store token in localStorage (or secure cookie)
5. Store user info in auth store
6. Redirect to /dashboard
7. Add token to all API requests (Authorization: Bearer token)
8. On 401 error: Clear token → redirect to /login
```

### Pinia Stores Required

```typescript
// stores/auth.ts
{
  user: { id, username, hospital_name, role }
  token: string
  isAuthenticated: boolean
  isLoading: boolean
  
  actions:
  - login(username, password)
  - logout()
  - refreshUser()
}

// stores/my-requests.ts
{
  requests: Array<Request>
  selectedRequest: Request | null
  isLoading: boolean
  error: string | null
  pagination: { page, limit, total }
  filters: { status, dateRange, search }
  
  actions:
  - fetchRequests(filters, pagination)
  - fetchRequestById(id)
  - createRequest(data)
  - updateRequest(id, data)
  - cancelRequest(id, reason)
  - setFilters(filters)
}

// stores/my-profile.ts
{
  profile: Hospital
  isLoading: boolean
  error: string | null
  
  actions:
  - fetchProfile()
  - updateProfile(data)
  - changePassword(oldPwd, newPwd)
}

// stores/dashboard.ts
{
  summary: { totalRequests, pending, completed, cancelled }
  charts: { statusDistribution, requestsOverTime }
  isLoading: boolean
  error: string | null
  
  actions:
  - fetchDashboard()
}
```

---

## 🔌 API INTEGRATION

### API Services to Create

```typescript
// api/auth.ts
- login(username: string, password: string)

// api/my-requests.ts
- getAll(filters, pagination)
- getById(id)
- create(data)
- update(id, data)
- delete(id)
- updateStatus(id, status, reason)

// api/my-profile.ts
- getProfile()
- updateProfile(data)
- changePassword(oldPwd, newPwd)

// api/dashboard.ts
- getStatusSummary()
```

### HTTP Client
```
✅ Use axios (from frontend-admin setup)
✅ Add Authorization header (Bearer token)
✅ Handle 401 → redirect to login
✅ Handle 403 → show "Access Denied"
✅ Add request/response interceptors
✅ Handle network errors gracefully
```

---

## 🔄 REAL-TIME FEATURES (WebSocket)

```
Connection:
├─ Endpoint: GET /ws/connect?token=JWT_TOKEN
├─ Connect on app load
├─ Maintain connection (auto-reconnect)
└─ Show connection indicator

Message Types:
├─ "request_status_changed": Status update notification
├─ "notification": General notification
└─ "alert": Urgent alert

UI Updates:
├─ Toast notification on status change
├─ Update request in list if open
├─ Refresh dashboard stats
├─ Update status color immediately

Connection Indicator:
├─ "Online" (green dot) in header
├─ "Offline" (gray dot) when disconnected
└─ Auto-reconnect attempts
```

---

## 📊 COMPONENTS & UI ELEMENTS

### Reusable from frontend-admin
```
Layout:
- AppLayout
- AppHeader
- AppSidebar
- AppFooter

Forms:
- TextInput
- SelectInput
- DateInput
- RadioInput
- TextArea
- FileUpload
- FormError
- FormLabel

Data Display:
- Table (with sorting, pagination)
- Card
- Badge
- Chip
- Avatar

Feedback:
- Modal/Dialog
- Toast/Notification
- Loading Skeleton
- Empty State
- Error Message

Navigation:
- Button (primary, secondary, danger)
- Link
- Breadcrumb
```

### New Components for RS Portal
```
- RequestStatusTimeline (status change history)
- BloodTypeSelector (ABO + Rhesus combined)
- ComponentRequestForm (dynamic component rows)
- RequestCard (compact request summary)
- HospitalBranding (header element)
```

---

## 📱 RESPONSIVE DESIGN

### Breakpoints
```
Mobile (< 640px):
├─ Single column
├─ Stack navigation
├─ Full-width forms
└─ Horizontal scroll tables

Tablet (640px - 1024px):
├─ Sidebar collapsible
├─ 2-column layouts
└─ Simplified charts

Desktop (> 1024px):
├─ Full sidebar
├─ Multi-column layouts
└─ Full-featured charts
```

### Mobile Features
```
✅ Bottom navigation tab bar (alternative to sidebar)
✅ Hamburger menu (collapsible sidebar)
✅ Touch-friendly buttons (min 44x44px)
✅ Readable fonts (min 16px)
✅ Optimized input types (number keypad for quantities)
✅ Full-screen modals
```

---

## 🎯 KEY FEATURES & BEHAVIORS

### Request Status Workflow
```
Status Lifecycle:
dibuat
  ├─ (RS creates)
  └─ (Admin can process)
      ↓
diproses
  ├─ (Admin reviewing)
  └─ (Admin can mark as ready)
      ↓
bisa_diambil
  ├─ (Ready for pickup)
  ├─ (RS can confirm pickup)
  └─ (Admin can mark as delivered)
      ↓
selesai
  └─ (Completed)

OR at any point:
dibatalkan
  ├─ (RS or Admin can cancel)
  └─ (With reason required)
```

### RS Actions by Status
```
dibuat       → Can Edit, Can Cancel, View
diproses     → Can Cancel (with reason), View
bisa_diambil → Can Confirm Pickup, Can Cancel, View
selesai      → Can View Only
dibatalkan   → Can View Only
```

### Data Filtering & Search
```
My Requests Page:
├─ Filter by Status (dropdown)
├─ Filter by Date Range (from-to)
├─ Search (Patient Name or Request ID)
└─ Sort by: Date (desc), Status, Patient Name
```

### Validation & Error Handling
```
Form Validation:
├─ Client-side: Real-time feedback
├─ Server-side: Show API errors
└─ Display clear error messages

Network Errors:
├─ Show toast notification
├─ Offer retry button
└─ Show fallback message

Auth Errors:
├─ 401: Auto redirect to login
├─ 403: Show "Access Denied"
└─ Clear stored token on 401
```

### Loading & Skeleton States
```
Page Loading:
├─ Show skeleton while fetching
└─ Fade in content when ready

Form Submission:
├─ Disable button
├─ Show loading spinner
└─ Prevent double submission
```

---

## 📋 ACCESSIBILITY

### Standards
- WCAG 2.1 Level AA compliance
- Semantic HTML
- Keyboard navigation (Tab, Enter, Esc)
- ARIA labels & roles
- Color contrast (min 4.5:1 for text)
- Focus indicators visible
- Form labels associated

### Implementation
```
✅ Use semantic tags (button, form, input, label)
✅ Add aria-labels for complex elements
✅ Ensure keyboard access to all interactive elements
✅ Test with screen readers
✅ High contrast for text
✅ Skip to main content link
```

---

## 🧪 TESTING REQUIREMENTS

### Unit Tests
- Store actions
- Utility functions
- Component logic

### Integration Tests
- API calls (mock)
- Store updates
- Navigation & routing

### E2E Tests
- Login flow
- Create request flow
- View requests flow
- Edit request flow
- Cancel request flow
- Profile flow

---

## 📦 PROJECT STRUCTURE

```
frontend-user/
├── src/
│   ├── api/              # API services
│   │   ├── auth.ts
│   │   ├── my-requests.ts
│   │   ├── my-profile.ts
│   │   ├── dashboard.ts
│   │   └── client.ts
│   │
│   ├── components/       # Vue components
│   │   ├── layout/
│   │   ├── common/
│   │   └── rs/
│   │
│   ├── stores/          # Pinia stores
│   │   ├── auth.ts
│   │   ├── my-requests.ts
│   │   ├── my-profile.ts
│   │   └── dashboard.ts
│   │
│   ├── views/           # Page components
│   │   ├── auth/
│   │   ├── dashboard/
│   │   ├── requests/
│   │   └── profile/
│   │
│   ├── types/           # TypeScript types
│   ├── router/          # Route config
│   ├── styles/          # CSS/Tailwind
│   ├── utils/           # Helper functions
│   ├── App.vue
│   └── main.ts
│
└── package.json
```

---

## 🚀 IMPLEMENTATION PHASES

### Phase 1: Setup (Days 1-2)
- Project scaffolding
- API client setup
- Pinia stores
- Router setup
- Base layout

### Phase 2: Authentication (Days 2-3)
- Login page
- Auth store
- Route protection
- Token management

### Phase 3: Dashboard (Days 3-4)
- Dashboard page
- Widgets/cards
- Charts

### Phase 4: My Requests List (Days 4-5)
- List page
- Table
- Filtering/search
- Pagination

### Phase 5: Request Detail (Days 5-6)
- Detail page
- Status timeline
- Actions

### Phase 6: Create/Edit Request (Days 6-8)
- Form page
- Validation
- Dynamic components

### Phase 7: Profile (Days 8-9)
- Profile page
- Edit form
- Password change

### Phase 8: Real-time (Days 9-10)
- WebSocket
- Notifications
- Status updates

### Phase 9-10: Testing, Polish, Deploy (Days 10-14)
- Unit tests
- Integration tests
- E2E tests
- Performance optimization
- Deployment

---

## 💡 IMPLEMENTATION TIPS

### Code Quality
- Group by feature, not type
- Keep components small & focused
- Reuse components from frontend-admin
- Use TypeScript for type safety

### Performance
- Lazy load pages (route code splitting)
- Minimize bundle size
- Cache appropriately
- Debounce search/filter

### Security
- Store token securely (httpOnly cookie)
- Validate all user inputs
- Don't expose sensitive data
- Use HTTPS only
- Sanitize user inputs

### Maintainability
- Write clear comments
- Follow Vue 3 patterns
- Keep functions pure
- Use meaningful commits

---

## ✅ PRE-DEVELOPMENT CHECKLIST

Before starting:
- [ ] Backend APIs tested & stable
- [ ] Token format confirmed
- [ ] Response format standardized
- [ ] Error handling strategy defined
- [ ] WebSocket confirmed working
- [ ] Design mockups approved
- [ ] Testing strategy planned
- [ ] Deployment pipeline ready

---

## 📚 REFERENCES

- Backend API Docs: `/backend/api-documentation/`
- Frontend-Admin: `/frontend-admin/`
- Vue 3: https://vuejs.org/
- TypeScript: https://www.typescriptlang.org/
- Tailwind: https://tailwindcss.com/
- Pinia: https://pinia.vuejs.org/

---

**Version**: 1.0  
**Status**: Ready for Implementation  
**Last Updated**: 2026-04-28
