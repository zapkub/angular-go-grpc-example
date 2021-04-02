export * from './bookstore.service';
import { BookstoreService } from './bookstore.service';
export * from './reportGenerator.service';
import { ReportGeneratorService } from './reportGenerator.service';
export const APIS = [BookstoreService, ReportGeneratorService];
