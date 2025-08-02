# Git clone stage
FROM rust:1.87-slim AS source
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*
ARG GIT_REPO=https://github.com/HERAFoundation/aip.git
ARG GIT_BRANCH=main
WORKDIR /source
RUN git clone --branch ${GIT_BRANCH} --single-branch --depth 1 ${GIT_REPO} .

# Build stage
FROM rust:1.87-slim AS builder

# Install required system dependencies for building
RUN apt-get update && apt-get install -y \
    pkg-config \
    libssl-dev \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy source code from git clone stage
COPY --from=source /source .

ARG FEATURES=embed,postgres
ARG GIT_HASH=0
ENV GIT_HASH=$GIT_HASH
ENV HTTP_TEMPLATE_PATH=/app/templates/

# Build the actual application with embed feature only
RUN cargo build --release --bin aip --no-default-features --features ${FEATURES}
# Build the client management CLI
RUN cargo build --release --bin aip-client-management --no-default-features --features ${FEATURES}
# Add the sqlx cli for running migrations in containers
RUN cargo install sqlx-cli --root /app/.cargo --no-default-features --features native-tls,postgres

# Runtime stage using distroless
FROM gcr.io/distroless/cc-debian12

# Add OCI labels
LABEL org.opencontainers.image.title="aip"
LABEL org.opencontainers.image.description="ATProtocol Identity Provider - OAuth 2.1 authorization server with ATProtocol integration"
LABEL org.opencontainers.image.version="0.1.0"
LABEL org.opencontainers.image.authors="Graze Social"
LABEL org.opencontainers.image.licenses="MIT"

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/target/release/aip /app/aip
# Copy the client management binary from builder stage
COPY --from=builder /app/target/release/aip-client-management /app/aip-client-management

# Copy static directory
COPY --from=builder /app/static ./static
# Copy migrations directory
COPY --from=builder /app/migrations ./migrations
# Copy sqlx binary for running migrations
COPY --from=builder /app/.cargo/bin/sqlx /app/sqlx

# Set default environment variables
ENV HTTP_STATIC_PATH=/app/static
ENV HTTP_PORT=8080

# Expose port
EXPOSE 8080

# Run the application
CMD ["/app/aip"]
