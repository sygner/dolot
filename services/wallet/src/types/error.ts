import * as grpc from '@grpc/grpc-js';

export class CustomError extends Error {
    public code: number;
    public details?: string;

    constructor(message: string, code: number, details?: string) {
        const defaultMessage = CustomError.getDefaultMessage(code);
        super(message || defaultMessage);

        this.name = this.constructor.name;
        this.code = code;
        this.details = details;

        Error.captureStackTrace(this, this.constructor);
    }

    public toGrpcStatus(): grpc.StatusObject {
        return {
            code: this.toGrpcCode(this.code),
            details: this.message,
            metadata: this.details ? this.createMetadata() : undefined as any, // Explicitly cast undefined
        };
    }
    
    public toHttpError(): { statusCode: number; message: string; details?: string } {
        return {
            statusCode: this.code, // Maps to HTTP status code
            message: this.message,
            details: this.details,
        };
    }

    private toGrpcCode(httpCode: number): grpc.status {
        const map: Record<number, grpc.status> = {
            400: grpc.status.ABORTED,
            401: grpc.status.UNAUTHENTICATED,
            403: grpc.status.PERMISSION_DENIED,
            404: grpc.status.NOT_FOUND,
            409: grpc.status.ALREADY_EXISTS,
            500: grpc.status.INTERNAL,
        };
        return map[httpCode] || grpc.status.UNKNOWN;
    }

    private static getDefaultMessage(code: number): string {
        const defaultMessages: Record<number, string> = {
            200: 'OK',
            201: 'Created',
            400: 'Bad Request',
            403: 'Permission Denied',
            404: 'Not Found',
            500: 'Internal Server Error',
        };
        return defaultMessages[code] || 'Unknown Error';
    }

    private createMetadata(): grpc.Metadata {
        const metadata = new grpc.Metadata();
        metadata.add('error-details', this.details!);
        return metadata;
    }
}
