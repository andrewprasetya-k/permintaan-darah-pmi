# Frontend-User Design Guidelines

## Quick Start

Dua file panduan utama telah dibuat:

1. **DESIGN_GUIDELINES.md** - Panduan desain lengkap
   - Overview aplikasi
   - Layout & design system
   - Detail setiap page (6 pages)
   - State management (Pinia stores)
   - API integration
   - Real-time features (WebSocket)
   - Components & responsive design
   - Testing requirements
   - Project structure
   - Implementation phases (10 phases)

2. **IMPLEMENTATION_CHECKLIST.md** - Checklist implementasi
   - Setup phase
   - Authentication
   - Dashboard
   - My Requests (list & detail)
   - Create/Edit Request
   - Profile Management
   - Real-time Features
   - UI Components
   - Testing
   - Responsive Design
   - Accessibility
   - Performance
   - Security
   - Deployment

---

## 📊 Pages Overview

| # | Page | Purpose | Route | Key Features |
|---|------|---------|-------|--------------|
| 1 | Login | Authentication | `/login` | Form validation, token storage, rate limiting |
| 2 | Dashboard | Overview & stats | `/` | Cards, charts, real-time updates |
| 3 | My Requests | List all requests | `/requests` | Pagination, filtering, search, sorting |
| 4 | Request Detail | View full details | `/requests/:id` | Status timeline, actions, history |
| 5 | Create/Edit | Request form | `/requests/new`, `/requests/:id/edit` | Multi-section form, validation, components |
| 6 | Profile | Hospital settings | `/profile` | Edit info, change password, account settings |

---

## 🔑 Backend Endpoints Required

```
Auth:
POST   /auth/login/rumah-sakit

Hospital Profile:
GET    /rumah-sakit/me
PUT    /rumah-sakit/me

My Requests:
GET    /permintaan-darah/my-requests
POST   /permintaan-darah
GET    /permintaan-darah/:id
PUT    /permintaan-darah/my-requests/:id
PUT    /permintaan-darah/update/:id
DELETE /permintaan-darah/:id

Dashboard:
GET    /dashboard/status-summary/{hospitalId}

WebSocket:
GET    /ws/connect?token=JWT_TOKEN
```

---

## 🎨 Design System

**Layout**: Header + Sidebar (3-4 menu items) + Main + Footer  
**Colors**: Blue (primary), Green (success), Red (danger), Yellow (warning)  
**Status Colors**:
- dibuat → Gray
- diproses → Blue
- bisa_diambil → Green
- selesai → Dark Green
- dibatalkan → Red

**Reuse from frontend-admin**: ~85% (layout, components, design)  
**Build new**: ~30% (RS-specific pages & logic)

---

## 📦 Pinia Stores

```
1. auth.ts - Login, user info, token
2. my-requests.ts - Requests list, detail, filters
3. my-profile.ts - Hospital profile
4. dashboard.ts - Dashboard stats & charts
```

---

## 🧪 Testing

- Unit tests (store actions, utilities)
- Integration tests (API mocking, routing)
- E2E tests (login, create, edit, profile flows)

---

## ⏱️ Estimated Timeline

- **Phase 1-2** (Days 1-3): Setup, Auth
- **Phase 3-5** (Days 3-6): Dashboard, My Requests List & Detail
- **Phase 6-7** (Days 6-9): Create/Edit Request, Profile
- **Phase 8-10** (Days 9-14): Real-time, Testing, Polish, Deploy

**Total: 14-20 days**

---

## 🚀 Getting Started

1. **Read** DESIGN_GUIDELINES.md for complete design specifications
2. **Reference** IMPLEMENTATION_CHECKLIST.md for task breakdown
3. **Follow** the implementation phases in order
4. **Reuse** ~85% from frontend-admin (components, stores, design)
5. **Test** each phase before moving to next

---

## 💡 Key Principles

✅ Reuse from frontend-admin (same tech stack, design language)  
✅ Keep RS portal simple & focused (3-4 pages vs 7 admin pages)  
✅ Own data only (RS can only see/manage their own requests)  
✅ Real-time updates via WebSocket  
✅ Mobile-first responsive design  
✅ Accessibility (WCAG 2.1 Level AA)  
✅ Type-safe with TypeScript  
✅ Comprehensive testing

---

## 📚 Files Created

- `DESIGN_GUIDELINES.md` - Complete design specifications
- `IMPLEMENTATION_CHECKLIST.md` - Detailed checklist
- `README_DESIGN.md` - This file (quick reference)

---

**Created**: 2026-04-28  
**Version**: 1.0  
**Status**: Ready for Implementation
