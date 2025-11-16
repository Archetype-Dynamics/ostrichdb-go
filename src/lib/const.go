package lib
/*
 *  Author: Marshall A Burns
 *  GitHub: @SchoolyB
 *
 *  Copyright (c) 2025-Present Archetype Dynamics, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

const OSTRICHDB_ADDRESS = "http://localhost:8042/api/v1"
const (
	NULL RecordType = iota
	CHAR
	STRING
	INTEGER
	FLOAT
	BOOLEAN
	DATE
	TIME
	DATETIME
	UUID
	CHAR_ARRAY //[]CHAR
	STRING_ARRAY //[]STRING
	INTEGER_ARRAY //[]INTEGER
	FLOAT_ARRAY //[]FLOAT
	BOOLEAN_ARRAY //[]BOOLEAN
	DATE_ARRAY //[]DATE
	TIME_ARRAY //[]TIME
	DATETIME_ARRAY //[]DATETIME
	UUID_ARRAY //[]UUID
)
