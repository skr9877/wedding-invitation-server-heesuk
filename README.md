# 모바일 청첩장 서버

이 프로젝트는 [모바일 청첩장](https://github.com/heesuk/wedding-invitation) 웹 애플리케이션의 백엔드 서버입니다. 모바일 청첩장에 필요한 API 엔드포인트와 데이터베이스 관리 기능을 제공합니다. 모바일 청첩장에 필요한 간단한 기능만 구현하였으며, 트래픽이 많지 않은 환경이기에 SQLite를 사용합니다.

SQLite 드라이버로 [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)(순수 Go 구현)를 사용하므로, gcc 등 별도의 C 컴파일러나 CGO 설정 없이 `go build`/`go run`만으로 바로 빌드·실행됩니다.

## 사전 요구사항

- Go 1.25 이상 (go.mod 기준. `go version`으로 확인)

## 제공 기능

- 방명록 작성 및 조회 API
  - 관리자 비밀번호를 통한 방명록 강제 삭제 기능
- 참석 의사 전달 API
  - 참석자 조회 기능은 현재 미구현 상태

## 시작하기

1. 저장소 복제:

   ```bash
   git clone https://github.com/heesuk/wedding-invitation-server.git
   cd wedding-invitation-server
   ```

2. 의존성 설치:

   ```bash
   go mod download
   ```

3. 환경변수 설정:

   환경변수 샘플은 `.env.sample` 파일에 저장되어 있습니다. 이 파일을 복사하여 `.env` 파일을 생성하고 각 환경변수를 수정합니다. `.env` 파일이 없으면 서버가 시작 시 바로 종료(panic)되므로 반드시 만들어야 합니다.

   ```bash
   cp .env.sample .env
   ```

   - `ALLOW_ORIGIN`
     - 허용할 프론트엔드 도메인 (예: `http://localhost:3000`). CORS 설정에 사용되며, 여기 등록되지 않은 origin에서의 요청은 브라우저에서 차단됩니다.
   - `ADMIN_PASSWORD`
     - 관리자 전용 비밀번호
     - 방명록 강제 삭제를 원하는 경우 해당 비밀번호로 삭제 가능

4. 서버 실행:
   ```bash
   go run .
   ```

   서버가 기본적으로 `http://localhost:8080`에서 실행되며, 프로젝트 루트에 `sql.db` 파일(SQLite)이 자동으로 생성됩니다.

5. 정상 동작 확인:
   ```bash
   curl "http://localhost:8080/api/guestbook?offset=0&limit=10"
   # {"posts":[],"total":0} 가 반환되면 정상
   ```

## 배포하기

1. 프로젝트 빌드:
   ```bash
   go build -o wedding-invitation-server .
   ```

2. 빌드된 바이너리 실행:
   ```bash
   ./wedding-invitation-server
   ```

   실행 파일과 같은 위치에 `.env` 파일이 있어야 하며, `sql.db`가 없으면 최초 실행 시 자동 생성됩니다.

## API 요약

| Method | Path | 설명 |
| --- | --- | --- |
| GET | `/api/guestbook?offset=&limit=` | 방명록 목록 조회 (`valid=true`인 글만, 최신순) |
| POST | `/api/guestbook` | 방명록 작성 `{name, content, password}` |
| PUT | `/api/guestbook` | 방명록 삭제 `{id, password}` (작성 시 비밀번호 또는 `ADMIN_PASSWORD`) |
| POST | `/api/attendance` | 참석 의사 등록 `{side, name, meal, count}` |
