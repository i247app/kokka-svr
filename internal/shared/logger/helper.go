package logger

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// extractIPAddress extracts the client's IP address from the request
func extractIPAddress(r *http.Request) string {
	if r == nil {
		return "0.0.0.0"
	}

	// Try X-Forwarded-For header first (for proxied requests)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		if before, _, ok := strings.Cut(xff, ","); ok {
			return strings.TrimSpace(before)
		}
		return strings.TrimSpace(xff)
	}

	// Try X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	ipAddr := r.RemoteAddr
	// Remove port if present
	if idx := strings.LastIndex(ipAddr, ":"); idx != -1 {
		ipAddr = ipAddr[:idx]
	}
	return ipAddr
}

// extractSessionInfoFromRequest extracts token, userid, route, and IP from the request
func extractSessionInfoFromRequest(r *http.Request) (token string, userid string, route string, ip string) {
	token = "anon"
	userid = "anon"
	route = "GET /unknown"
	ip = "0.0.0.0"

	if r == nil {
		return
	}

	// // Get session from request context
	// sess := session.GetRequestSession(r)
	// if sess == nil {
	// 	return
	// }

	// // Extract token (key) from session
	// if key, ok := sess.Get("key"); ok {
	// 	if keyStr, ok := key.(string); ok && len(keyStr) >= 6 {
	// 		if len(keyStr) >= 6 {
	// 			token = keyStr[len(keyStr)-6:]
	// 		} else {
	// 			token = keyStr
	// 		}
	// 	}
	// }

	// // Extract userid from session
	// if uid, ok := sess.UID(); ok {
	// 	userid = uid
	// }

	// Extract route from context if available
	route = fmt.Sprintf("%s %s", r.Method, r.URL.Path)

	// Extract IP address
	ip = extractIPAddress(r)

	return
}

// withSessionInfo adds session info (token, userid, route, IP) to context for logging
func withSessionInfo(ctx context.Context, token, userid, route, ipAddr string) context.Context {
	ctx = context.WithValue(ctx, tokenKey, token)
	ctx = context.WithValue(ctx, useridKey, userid)
	ctx = context.WithValue(ctx, routeKey, route)
	ctx = context.WithValue(ctx, ipKey, ipAddr)
	return ctx
}

// extractSessionInfo retrieves token, userid, route, and IP from context
func extractSessionInfo(ctx context.Context) (token, userid, route, ip string) {
	token = "anon"
	userid = "anon"
	ip = "0.0.0.0"

	if t, ok := ctx.Value(tokenKey).(string); ok && t != "" {
		token = t
	}
	if u, ok := ctx.Value(useridKey).(string); ok && u != "" {
		userid = u
	}
	if r, ok := ctx.Value(routeKey).(string); ok && r != "" {
		route = r
	}
	if i, ok := ctx.Value(ipKey).(string); ok && i != "" {
		ip = i
	}

	return
}
